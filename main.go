package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatalf("token not found")
	}
}
