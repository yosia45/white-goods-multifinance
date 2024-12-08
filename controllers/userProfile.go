package controllers

import (
	"white-goods-multifinace/repositories"

	"github.com/labstack/echo/v4"
)

type UserProfileController struct {
	userProfileRepo repositories.UserProfileRepository
}

func NewUserProfileController(userProfileRepo repositories.UserProfileRepository) *UserProfileController {
	return &UserProfileController{
		userProfileRepo: userProfileRepo,
	}
}

func (upc *UserProfileController) UpdateUserProfile(c echo.Context) error {
	return nil
}
