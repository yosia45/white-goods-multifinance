package controllers

import (
	"net/http"
	"white-goods-multifinace/dto"
	"white-goods-multifinace/models"
	"white-goods-multifinace/repositories"
	"white-goods-multifinace/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserLimitController struct {
	userLimit repositories.UserLimitRepository
}

func NewUserLimitController(userLimit repositories.UserLimitRepository) *UserLimitController {
	return &UserLimitController{
		userLimit: userLimit,
	}
}

func (ulc *UserLimitController) CreateUserLimit(c echo.Context) error {
	var userLimitBody []dto.AddUserLimitBody
	var userLimits []models.UserLimit

	if err := c.Bind(&userLimitBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	for _, limit := range userLimitBody {

		if limit.UserID == "" {
			return utils.HandlerError(c, utils.NewBadRequestError("User ID is required"))
		}

		if limit.TenorID == "" {
			return utils.HandlerError(c, utils.NewBadRequestError("Tenor ID is required"))
		}

		parsedUserID, err := uuid.Parse(limit.UserID)
		if err != nil {
			return utils.HandlerError(c, utils.NewBadRequestError("Invalid user ID"))
		}

		parsedTenorID, err := uuid.Parse(limit.TenorID)
		if err != nil {
			return utils.HandlerError(c, utils.NewBadRequestError("Invalid tenor ID"))
		}

		userLimits = append(userLimits, models.UserLimit{
			UserID:         parsedUserID,
			TenorID:        parsedTenorID,
			Limit:          limit.Limit,
			CurrentBalance: limit.Limit,
		})
	}

	if err := ulc.userLimit.CreateUserLimit(&userLimits); err != nil {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User limit created successfully"})
}
