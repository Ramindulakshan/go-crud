package main

import (
	"log"
	"net/http"
	"os"

	httpRouter "go-crud/internal/http"
	"go-crud/internal/user"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	repo := user.NewInMemoryRepository()
	svc := user.NewService(repo)
	handler := user.NewHandler(svc)
	router := httpRouter.NewRouter(handler)

	addr := ":" + port
	log.Printf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
