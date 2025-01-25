package handlers

import (
	"UserSystem/models"
	"UserSystem/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.service.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user.ID = id

	if err := h.service.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	if err := h.service.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
