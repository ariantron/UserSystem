package repositories

import (
	"UserSystem/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id uuid.UUID) (models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Preload("Addresses").Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.DB.Preload("Addresses").First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *userRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.User{}, "id = ?", id).Error
}
