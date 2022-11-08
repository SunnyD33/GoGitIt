package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvWithKey(key string) string {
	return os.Getenv(key)
}

func loadEnv() {
	err := godotenv.Load("../../env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}