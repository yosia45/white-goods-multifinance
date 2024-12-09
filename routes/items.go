package routes

import (
	"white-goods-multifinace/configs"
	"white-goods-multifinace/controllers"
	"white-goods-multifinace/repositories"

	"github.com/labstack/echo/v4"
)

func ItemRoutes(e *echo.Echo) {
	itemRepo := repositories.NewItemRepository(configs.DB)
	itemTenorRepo := repositories.NewItemTenorRepository(configs.DB)
	itemController := controllers.NewItemController(itemRepo, itemTenorRepo)

	e.POST("/items", itemController.CreateItem)
}
