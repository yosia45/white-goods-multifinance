package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORSConfig() echo.MiddlewareFunc {
	origins := []string{os.Getenv("PRODUCTION_CLIENT_URL")}
	if os.Getenv("APP_ENV") == "development" {
		origins = append(origins, os.Getenv("DEVELOPMENT_CLIENT_URL"))
	}

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization,
		},
		AllowMethods: []string{
			echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS,
		},
	})
}
