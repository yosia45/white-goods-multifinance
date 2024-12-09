package routes

import (
	"white-goods-multifinace/configs"
	"white-goods-multifinace/controllers"
	"white-goods-multifinace/middlewares"
	"white-goods-multifinace/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Echo) {
	transactionRepo := repositories.NewTransactionRepository(configs.DB)
	itemRepo := repositories.NewItemRepository(configs.DB)
	purchaseRepo := repositories.NewPurchaseRepository(configs.DB)
	userLimitRepo := repositories.NewUserLimitRepository(configs.DB)

	transactionController := controllers.NewTransactionController(transactionRepo, purchaseRepo, itemRepo, userLimitRepo)

	e.POST("/transactions", transactionController.CreateTransaction, middlewares.JWTAuth)
}
