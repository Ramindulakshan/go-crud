package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	dbMigrate "go-crud/internal/db"
	db "go-crud/internal/db/sqlc"
	httpRouter "go-crud/internal/http"
	"go-crud/internal/user"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	port := getEnv("APP_PORT", "8080")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "usersdb"),
		getEnv("DB_SSLMODE", "disable"),
	)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}

	if err := dbMigrate.Migrate(ctx, sqlDB); err != nil {
		log.Fatalf("failed to run migration: %v", err)
	}

	queries := db.New(sqlDB)
	repo := user.NewPostgresRepository(queries)
	svc := user.NewService(repo)
	handler := user.NewHandler(svc)
	router := httpRouter.NewRouter(handler)

	addr := ":" + port
	log.Printf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
