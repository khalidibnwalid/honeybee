package main

import (
	"khalidibnwalid/luma_server/internal/server"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := 8080

	if envPort, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		port = envPort
	}

	server := server.NewServer(port)
	log.Printf("Starting server on port %d...\n", port)
	server.ListenAndServe()
}
