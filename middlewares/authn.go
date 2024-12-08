package middlewares

import (
	"fmt"
	"os"
	"strings"
	"time"
	"white-goods-multifinace/dto"
	"white-goods-multifinace/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(userID uuid.UUID, role string, tokenDuration int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(tokenDuration * int(time.Minute))).Unix()

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.HandlerError(c, utils.NewForbiddenError("please login first"))
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return utils.HandlerError(c, utils.NewForbiddenError("invalid authorization header format"))
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			return utils.HandlerError(c, utils.NewUnauthorizedError("invalid token"))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() > int64(exp) {
					return utils.HandlerError(c, utils.NewUnauthorizedError("token has expired"))
				}
			} else {
				return utils.HandlerError(c, utils.NewUnauthorizedError("invalid token: no expiration time"))
			}

			userIDStr, ok := claims["user_id"].(string)
			if !ok {
				return utils.HandlerError(c, utils.NewUnauthorizedError("invalid token: user_id not found"))
			}

			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				return utils.HandlerError(c, utils.NewUnauthorizedError("invalid token: invalid user_id"))
			}

			role := claims["role"].(string)

			c.Set("userPayload", &dto.JWTPayload{
				UserID: userID,
				Role:   role,
			})
		} else {
			return utils.HandlerError(c, utils.NewUnauthorizedError("invalid token"))
		}

		return next(c)
	}
}
