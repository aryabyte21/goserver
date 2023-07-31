package main

import (
	"goly/model"
	"goly/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	 err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
	model.Setup()
	server.SetupAndListen()
}
