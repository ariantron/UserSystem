package main

import (
	"UserSystem/configs"
	"UserSystem/internal/handlers"
	"UserSystem/internal/repositories"
	"UserSystem/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Serve(db *gorm.DB) {
	e := echo.New()
	Router(e, db)
	e.Logger.Fatal(e.Start(":" + configs.AppPort))
}

func Router(e *echo.Echo, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	addressRepo := repositories.NewAddressRepository(db)

	userService := services.NewUserService(userRepo)
	addressService := services.NewAddressService(addressRepo)

	userHandler := handlers.NewUserHandler(userService)
	addressHandler := handlers.NewAddressHandler(addressService)

	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	e.POST("/addresses", addressHandler.CreateAddress)
	e.GET("/users/:userID/addresses", addressHandler.GetAddressesByUser)
	e.PUT("/addresses/:id", addressHandler.UpdateAddress)
	e.DELETE("/addresses/:id", addressHandler.DeleteAddress)
}
