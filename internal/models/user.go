package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Email       string    `gorm:"type:varchar(100);not null" json:"email"`
	PhoneNumber string    `gorm:"type:varchar(20);not null" json:"phone_number"`
	Addresses   []Address `gorm:"foreignKey:UserID" json:"addresses"`
}
