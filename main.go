package main

import (
	"gomongo/repository"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error read env file with err: %s", errEnv)
	}
	// repository.Insert()
	repository.Find()
}
