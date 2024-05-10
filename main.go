package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grodier/bootdev-rss-aggregator/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthcheck", handlers.Healthcheck)
	mux.HandleFunc("GET /v1/err", handlers.Err)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Running server on port %v\n", port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal("Issue with server:", err)
	}
}
