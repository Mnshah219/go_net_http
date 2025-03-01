package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	initializeLogging()
	router := registerRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	portString := fmt.Sprintf(":%s", port)
	slog.Info("Starting servet", "port", portString)
	err := http.ListenAndServe(portString, router)
	if err != nil {
		log.Fatal(err)
	}
}
