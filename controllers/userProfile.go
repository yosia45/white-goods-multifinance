package controllers

import (
	"fmt"
	"net/http"
	"strconv"
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
	userPayload := c.Get("userPayload").(*dto.JWTPayload)

	updatedUserBody.FullName = c.FormValue("full_name")
	if updatedUserBody.FullName == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Full name is required"))
	}

	updatedUserBody.LegalName = c.FormValue("legal_name")
	if updatedUserBody.LegalName == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Legal name is required"))
	}

	updatedUserBody.BirthPlace = c.FormValue("birth_place")
	if updatedUserBody.BirthPlace == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Birth place is required"))
	}

	birthDate := c.FormValue("birth_date")
	parsedBirthDate, err := time.Parse("2006-01-02T15:04:05Z", birthDate)
	if parsedBirthDate.IsZero() || err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Birth date is required"))
	}
	updatedUserBody.BirthDate = parsedBirthDate

	stringSalary := c.FormValue("salary")
	parsedSalary, err := strconv.ParseFloat(stringSalary, 64)
	if err != nil || parsedSalary == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Salary is required"))
	}

	ktpFile, err := c.FormFile("ktp_file")
	if err == nil {
		ktpFilePath, err := utils.SaveUploadFile(ktpFile, "assets/ktps")
		if err != nil {
			fmt.Println(err.Error(), "err SaveUploadFile")
			return utils.HandlerError(c, utils.NewInternalError("Failed to save KTP file"))
		}
		updatedUserBody.KTPFilePathURL = ktpFilePath
	} else {
		fmt.Println(err.Error())
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	selfieFile, err := c.FormFile("selfie_file")
	if err == nil {
		selfieFilePath, err := utils.SaveUploadFile(selfieFile, "assets/selfies")
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to save selfie file"))
		}
		updatedUserBody.SelfieFilePathURL = selfieFilePath
	} else {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	newUpdateProfile := models.UserProfile{
		LegalName:      &updatedUserBody.LegalName,
		BirthPlace:     &updatedUserBody.BirthPlace,
		BirthDate:      &updatedUserBody.BirthDate,
		Salary:         &updatedUserBody.Salary,
		KTPFilePath:    &updatedUserBody.KTPFilePathURL,
		SelfieFilePath: &updatedUserBody.SelfieFilePathURL,
	}

	if err := upc.userProfileRepo.UpdateUserProfile(&newUpdateProfile, userPayload.UserID); err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to update user profile"))
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User profile updated successfully",
	})
}
