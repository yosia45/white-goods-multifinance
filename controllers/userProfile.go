package controllers

import (
	"net/http"
	"white-goods-multifinace/dto"
	"white-goods-multifinace/models"
	"white-goods-multifinace/repositories"
	"white-goods-multifinace/utils"

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
	var updatedUserBody dto.UpdateUserProfileBody
	userPayload := c.Get("userPayload").(*dto.JWTPayload)
	var KTPFilePathURL, selfieFilePathURL string

	if updatedUserBody.FullName == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Full name is required"))
	}

	if updatedUserBody.LegalName == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Legal name is required"))
	}

	if updatedUserBody.BirthPlace == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Birth place is required"))
	}

	if updatedUserBody.BirthDate.IsZero() {
		return utils.HandlerError(c, utils.NewBadRequestError("Birth date is required"))
	}

	if updatedUserBody.Salary == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Salary is required"))
	}

	ktpFile, err := c.FormFile("ktp_file")
	if err == nil {
		ktpFilePath, err := utils.SaveUploadFile(ktpFile, "assets/ktps")
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to save KTP file"))
		}
		KTPFilePathURL = ktpFilePath
	} else {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	selfieFile, err := c.FormFile("selfie_file")
	if err == nil {
		selfieFilePath, err := utils.SaveUploadFile(selfieFile, "assets/selfies")
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to save selfie file"))
		}
		selfieFilePathURL = selfieFilePath
	} else {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	newUpdateProfile := models.UserProfile{
		LegalName:      &updatedUserBody.LegalName,
		BirthPlace:     &updatedUserBody.BirthPlace,
		BirthDate:      &updatedUserBody.BirthDate,
		Salary:         &updatedUserBody.Salary,
		KTPFilePath:    &KTPFilePathURL,
		SelfieFilePath: &selfieFilePathURL,
	}

	if err := upc.userProfileRepo.UpdateUserProfile(&newUpdateProfile, userPayload.UserID); err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to update user profile"))
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User profile updated successfully",
	})
}
