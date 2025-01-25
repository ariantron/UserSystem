package services

import (
	"UserSystem/models"
	"UserSystem/repositories"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uuid.UUID) (models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uuid.UUID) (models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.Delete(id)
}
