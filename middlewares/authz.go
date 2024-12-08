package middlewares

import (
	"white-goods-multifinace/dto"
	"white-goods-multifinace/utils"

	"github.com/labstack/echo/v4"
)

func Authz(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userPayload := c.Get("userPayload").(*dto.JWTPayload)

		if userPayload.Role != "admin" {
			return utils.HandlerError(c, utils.NewForbiddenError("only owner can access this resource"))
		}
		return next(c)
	}
}
