package repositories

import (
	"UserSystem/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address *models.Address) error
	FindByUserID(userID uuid.UUID) ([]models.Address, error)
	Update(address *models.Address) error
	Delete(id uuid.UUID) error
}

type addressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{DB: db}
}

func (r *addressRepository) Create(address *models.Address) error {
	return r.DB.Create(address).Error
}

func (r *addressRepository) FindByUserID(userID uuid.UUID) ([]models.Address, error) {
	var addresses []models.Address
	err := r.DB.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *addressRepository) Update(address *models.Address) error {
	return r.DB.Save(address).Error
}

func (r *addressRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Address{}, "id = ?", id).Error
}
