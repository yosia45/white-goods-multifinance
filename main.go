package main

import (
	"log"
	"os"
	"white-goods-multifinace/configs"

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

	// cli.Auth(e)
	// cli.RoomingHouseRoutes(e)
	// cli.SizeRoutes(e)
	// cli.AdditionalPriceRoutes(e)
	// cli.TransactionCategoryRoutes(e)
	// cli.PricingPackageRoutes(e)
	// cli.RoomRoutes(e)
	// cli.AdditionalPriceRoutes(e)
	// cli.TransactionRoutes(e)
	// cli.TenantRoutes(e)
	// cli.AdminRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
