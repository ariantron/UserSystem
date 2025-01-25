package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func BuildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DbHost,
		DbUser,
		DbPassword,
		DbName,
		DbPort,
		DbSslMode,
	)
}

func ConnectDB() *gorm.DB {
	dsn := BuildDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	return db
}
