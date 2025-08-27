package main

import (
	"log"
	"net/http"
	"os"

	exoplanetHandlers "github.com/almaraz333/go-astronomy-server/internal/handlers/exoplanets"
	tempHandlers "github.com/almaraz333/go-astronomy-server/internal/handlers/temp"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	http.HandleFunc("/api", tempHandlers.Test)
	http.HandleFunc("/api/exoplanets", exoplanetHandlers.GetExoplanets)

	log.Printf("Server running on %v", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
