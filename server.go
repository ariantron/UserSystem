package main

import (
	"UserSystem/handlers"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Serve(db *gorm.DB) {
	e := echo.New()
	Router(e, db)
	err := e.Start(":8080")
	if err != nil {
		fmt.Println("Application Serve has encountered error:\r\n" + err.Error())
	}
}

func Router(e *echo.Echo, db *gorm.DB) {
	userHandler := handlers.NewUserHandler(db)
	addressHandler := handlers.NewAddressHandler(db)

	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	// Address routes
	e.POST("/addresses", addressHandler.CreateAddress)
	e.GET("/users/:userID/addresses", addressHandler.GetAddressesByUser)
	e.PUT("/addresses/:id", addressHandler.UpdateAddress)
	e.DELETE("/addresses/:id", addressHandler.DeleteAddress)
}
