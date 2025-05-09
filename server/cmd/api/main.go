package main

import (
	"khalidibnwalid/luma_server/internal/handlers"
	"khalidibnwalid/luma_server/internal/server"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := 8080

	ctx, err := server.NewServerContext()
	if err != nil {
		log.Fatalf("Failed to create server context: %v", err)
	}

	if err := ctx.DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database")

	if envPort, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		port = envPort
	}

	server, err := handlers.NewServer(port)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	log.Printf("Starting server on port %d...\n", port)
	server.ListenAndServe()
}
