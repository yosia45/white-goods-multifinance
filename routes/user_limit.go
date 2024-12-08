package routes

import (
	"white-goods-multifinace/configs"
	"white-goods-multifinace/controllers"
	"white-goods-multifinace/middlewares"
	"white-goods-multifinace/repositories"

	"github.com/labstack/echo/v4"
)

func UserLimitRoutes(e *echo.Echo) {
	userLimitRepo := repositories.NewUserLimitRepository(configs.DB)
	userLimitController := controllers.NewUserLimitController(userLimitRepo)

	e.POST("/user-limit", userLimitController.CreateUserLimit, middlewares.JWTAuth, middlewares.Authz)
}
