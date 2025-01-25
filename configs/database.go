package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbSslMode  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	DbSslMode = os.Getenv("DB_SSL_MODE")

	if DbHost == "" || DbPort == "" || DbUser == "" || DbPassword == "" || DbName == "" || DbSslMode == "" {
		log.Fatalf("Missing required database environment variable(s) for DB configuration")
	}
}
