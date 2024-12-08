package routes

import (
	"white-goods-multifinace/configs"
	"white-goods-multifinace/controllers"
	"white-goods-multifinace/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	userRepo := repositories.NewUserRepository(configs.DB)
	userProfile := repositories.NewUserProfileRepository(configs.DB)
	userController := controllers.NewUserController(userRepo, userProfile)

	e.POST("/login", userController.Login)
	e.POST("/register-customer", userController.RegisterCustomer)
}
