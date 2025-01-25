package database

import (
	"UserSystem/configs"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func BuildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s statement_timeout=300000",
		configs.DbHost,
		configs.DbUser,
		configs.DbPassword,
		configs.DbName,
		configs.DbPort,
		configs.DbSslMode,
	)
}

func ConnectDB() *gorm.DB {
	dsn := BuildDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	var timeout string
	if err := db.Raw("SHOW statement_timeout").Scan(&timeout).Error; err != nil {
		log.Printf("Error retrieving statement_timeout: %v", err)
	} else {
		log.Printf("Current statement_timeout: %s", timeout)
	}

	return db
}
