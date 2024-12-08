package main

import (
	"log"
	"os"
	"white-goods-multifinace/configs"
	"white-goods-multifinace/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configs.InitDB()

	// seeders.SeedOTR(configs.DB)
	// seeders.SeedItems(configs.DB)
	// seeders.SeedTenor(configs.DB)

	port := os.Getenv("APP_DEVELOPMENT_PORT")

	e := echo.New()

	routes.AuthRoutes(e)
	// routes.UserProfileRoutes(e)
	routes.UserLimitRoutes(e)
	routes.PurchaseRoutes(e)
	routes.TransactionRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
