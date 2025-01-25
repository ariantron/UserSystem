package database

import (
	"UserSystem/internal/models"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Address{},
	)
	if err != nil {
		log.Fatalf("Error migrating the models: %v", err)
	}
	fmt.Println("Database migrated successfully")
}
