package controllers

import (
	"fmt"
	"net/http"
	"time"
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

	parsedTime, err := time.Parse("2006-01-02 15:04:05", updatedUserBody.BirthDate.Format("2006-01-02 15:04:05"))
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid birth date format"))
	}

	updatedUserBody.BirthDate = parsedTime

	userPayload := c.Get("userPayload").(*dto.JWTPayload)
	if err := c.Bind(&updatedUserBody); err != nil {
		fmt.Println(err)
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if updatedUserBody.LegalName == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Legal name is required"))
	}

	if updatedUserBody.FullName == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Full name is required"))
	}

	if updatedUserBody.NIK == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("NIK is required"))
	}

	if len(updatedUserBody.NIK) != 16 {
		return utils.HandlerError(c, utils.NewBadRequestError("NIK must be 16 characters long"))
	}

	if updatedUserBody.BirthPlace == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Birth place is required"))
	}

	// parsedBirthDate, err := time.Parse("2006-01-02 15:04:05", updatedUserBody.BirthDate)
	// if err != nil {
	// 	return utils.HandlerError(c, utils.NewBadRequestError("Invalid birth date format"))
	// }

	// if parsedBirthDate.IsZero() {
	// 	return utils.HandlerError(c, utils.NewBadRequestError("Birth date is required"))
	// }

	if updatedUserBody.Salary == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Salary is required"))
	}

	ktpFile, err := c.FormFile("ktp_file")
	if err != nil {
		ktpFilePath, err := utils.SaveUploadFile(ktpFile, "assets/ktps")
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to save KTP file"))
		}
		updatedUserBody.KTPFilePath = ktpFilePath
	}

	selfieFile, err := c.FormFile("selfie_file")
	if err != nil {
		selfieFilePath, err := utils.SaveUploadFile(selfieFile, "assets/selfies")
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to save selfie file"))
		}
		updatedUserBody.SelfieFilePath = selfieFilePath
	}

	newUpdateProfile := models.UserProfile{
		LegalName:      updatedUserBody.LegalName,
		FullName:       updatedUserBody.FullName,
		NIK:            updatedUserBody.NIK,
		BirthPlace:     updatedUserBody.BirthPlace,
		BirthDate:      &updatedUserBody.BirthDate,
		Salary:         updatedUserBody.Salary,
		KTPFilePath:    updatedUserBody.KTPFilePath,
		SelfieFilePath: updatedUserBody.SelfieFilePath,
	}

	if err := upc.userProfileRepo.UpdateUserProfile(&newUpdateProfile, userPayload.UserID); err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to update user profile"))
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User profile updated successfully",
	})
}
