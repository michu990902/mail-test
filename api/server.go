package api

import (
	"os"
	"log"

	"mailtest/api/controllers"

	env "github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run(){
	var err error
	err = env.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	server.Initialize(
		os.Getenv("SENDER_EMAIL"),
		os.Getenv("SENDER_PASS"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
	)
}