package routes

import (
	"white-goods-multifinace/configs"
	"white-goods-multifinace/controllers"
	"white-goods-multifinace/middlewares"
	"white-goods-multifinace/repositories"

	"github.com/labstack/echo/v4"
)

func PurchaseRoutes(e *echo.Echo) {
	purchaseRepo := repositories.NewPurchaseRepository(configs.DB)
	userRepo := repositories.NewUserRepository(configs.DB)
	itemRepo := repositories.NewItemRepository(configs.DB)
	itemTenorRepo := repositories.NewItemTenorRepository(configs.DB)
	userLimitRepo := repositories.NewUserLimitRepository(configs.DB)
	tenorRepo := repositories.NewTenorRepository(configs.DB)

	purchaseController := controllers.NewPurchaseController(userRepo, userLimitRepo, itemRepo, itemTenorRepo, purchaseRepo, tenorRepo)

	e.GET("/purchases", purchaseController.GetAllPurchase, middlewares.JWTAuth)
	e.GET("/purchases/:id", purchaseController.GetPurchaseByID, middlewares.JWTAuth)
	e.POST("/purchases", purchaseController.CreatePurchase, middlewares.JWTAuth)
}
