package routes

import (
	"white-goods-multifinace/configs"
	"white-goods-multifinace/controllers"
	"white-goods-multifinace/middlewares"
	"white-goods-multifinace/repositories"

	"github.com/labstack/echo/v4"
)

func UserProfileRoutes(e *echo.Echo) {
	userProfileRepo := repositories.NewUserProfileRepository(configs.DB)
	userProfileController := controllers.NewUserProfileController(userProfileRepo)

	e.PUT("/user-profile", userProfileController.UpdateUserProfile, middlewares.JWTAuth)
}
