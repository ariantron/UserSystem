package models

import "github.com/google/uuid"

type Address struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Street  string    `gorm:"type:varchar(100);not null" json:"street"`
	City    string    `gorm:"type:varchar(100);not null" json:"city"`
	State   string    `gorm:"type:varchar(100);not null" json:"state"`
	ZipCode string    `gorm:"type:varchar(100);not null" json:"zip_code"`
	Country string    `gorm:"type:varchar(100);not null" json:"country"`
	UserID  uuid.UUID `json:"user_id"`
}
