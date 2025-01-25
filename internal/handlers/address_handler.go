package handlers

import (
	"UserSystem/internal/models"
	"UserSystem/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AddressHandler struct {
	service services.AddressService
}

func NewAddressHandler(service services.AddressService) *AddressHandler {
	return &AddressHandler{service: service}
}

func (h *AddressHandler) CreateAddress(c echo.Context) error {
	address := new(models.Address)
	if err := c.Bind(address); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.service.CreateAddress(address); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, address)
}

func (h *AddressHandler) GetAddressesByUser(c echo.Context) error {
	userID, _ := uuid.Parse(c.Param("userID"))

	addresses, err := h.service.GetAddressesByUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, addresses)
}

func (h *AddressHandler) UpdateAddress(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	var address models.Address
	address.ID = id
	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.service.UpdateAddress(&address); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, address)
}

func (h *AddressHandler) DeleteAddress(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	if err := h.service.DeleteAddress(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
