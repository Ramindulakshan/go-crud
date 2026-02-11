package http_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	dbMigrate "go-crud/internal/db"
	db "go-crud/internal/db/sqlc"
	httprouter "go-crud/internal/http"
	"go-crud/internal/user"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func testDSN() string {
	if dsn := os.Getenv("TEST_DATABASE_URL"); dsn != "" {
		return dsn
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		return ""
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "usersdb"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		dbPass = "postgres"
	}
	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, dbUser, dbPass, dbName, sslMode)
}

func TestUsersAPIIntegration(t *testing.T) {
	dsn := testDSN()
	if dsn == "" {
		t.Skip("set TEST_DATABASE_URL or DB_* env vars to run integration tests")
	}

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer sqlDB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		t.Skipf("db unavailable for integration test: %v", err)
	}

	if err := dbMigrate.Migrate(ctx, sqlDB); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	if _, err := sqlDB.ExecContext(ctx, "TRUNCATE TABLE users"); err != nil {
		t.Fatalf("truncate users: %v", err)
	}

	queries := db.New(sqlDB)
	repo := user.NewPostgresRepository(queries)
	svc := user.NewService(repo)
	handler := user.NewHandler(svc)
	router := httprouter.NewRouter(handler)
	server := httptest.NewServer(router)
	defer server.Close()

	createPayload := map[string]any{
		"firstName": "John",
		"lastName":  "Doe",
		"email":     "john.integration@example.com",
		"phone":     "+14155552671",
		"status":    "Active",
	}
	b, _ := json.Marshal(createPayload)
	resp, err := http.Post(server.URL+"/users", "application/json", bytes.NewReader(b))
	if err != nil {
		t.Fatalf("post /users: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201 from create, got %d", resp.StatusCode)
	}

	var created map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		t.Fatalf("decode create response: %v", err)
	}
	id, _ := created["userId"].(string)
	if id == "" {
		t.Fatal("expected userId in create response")
	}

	getResp, err := http.Get(server.URL + "/users/" + id)
	if err != nil {
		t.Fatalf("get /users/{id}: %v", err)
	}
	defer getResp.Body.Close()
	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 from get by id, got %d", getResp.StatusCode)
	}
}
