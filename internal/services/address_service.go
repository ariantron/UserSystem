package services

import (
	"UserSystem/internal/models"
	"UserSystem/internal/repositories"
	"github.com/google/uuid"
)

type AddressService interface {
	CreateAddress(address *models.Address) error
	GetAddressesByUser(userID uuid.UUID) ([]models.Address, error)
	UpdateAddress(address *models.Address) error
	DeleteAddress(id uuid.UUID) error
}

type addressService struct {
	repo repositories.AddressRepository
}

func NewAddressService(repo repositories.AddressRepository) AddressService {
	return &addressService{repo: repo}
}

func (s *addressService) CreateAddress(address *models.Address) error {
	return s.repo.Create(address)
}

func (s *addressService) GetAddressesByUser(userID uuid.UUID) ([]models.Address, error) {
	return s.repo.FindByUserID(userID)
}

func (s *addressService) UpdateAddress(address *models.Address) error {
	return s.repo.Update(address)
}

func (s *addressService) DeleteAddress(id uuid.UUID) error {
	return s.repo.Delete(id)
}
