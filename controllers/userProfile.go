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

	userProfile, err := upc.userProfileRepo.FindUserProfileByUserID(userPayload.UserID)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to get user profile"))
	}

	ktpFile, err := c.FormFile("ktp_file")
	if err == nil {
		fileType, err := utils.GetFileTypeByExtension(ktpFile)
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to get file type"))
		}

		if fileType != "image" {
			return utils.HandlerError(c, utils.NewBadRequestError("KTP file must be an image type"))
		}

		if userProfile.KTPFileURL != "" {
			if err := utils.DeleteFile(userProfile.KTPFileURL); err != nil {
				fmt.Println(err)
				return utils.HandlerError(c, utils.NewInternalError("Failed to delete KTP file"))
			}
		}

		ktpFilePath, err := utils.SaveUploadFile(ktpFile, "assets/ktps")
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to save KTP file"))
		}
		updatedUserBody.KTPFilePathURL = ktpFilePath
	} else {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	selfieFile, err := c.FormFile("selfie_file")
	if err == nil {
		fileType, err := utils.GetFileTypeByExtension(selfieFile)
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to get file type"))
		}

		if fileType != "image" {
			return utils.HandlerError(c, utils.NewBadRequestError("Selfie file must be an image type"))
		}

		if userProfile.SelfieURL != "" {
			if err := utils.DeleteFile(userProfile.SelfieURL); err != nil {
				return utils.HandlerError(c, utils.NewInternalError("Failed to delete Selfie file"))
			}
		}

		selfieFilePath, err := utils.SaveUploadFile(selfieFile, "assets/selfies")
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to save selfie file"))
		}
		updatedUserBody.SelfieFilePathURL = selfieFilePath
	} else {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	newUpdateProfile := models.UserProfile{
		LegalName:      updatedUserBody.LegalName,
		BirthPlace:     updatedUserBody.BirthPlace,
		BirthDate:      &updatedUserBody.BirthDate,
		Salary:         parsedSalary,
		KTPFilePath:    updatedUserBody.KTPFilePathURL,
		SelfieFilePath: updatedUserBody.SelfieFilePathURL,
	}

	if err := upc.userProfileRepo.UpdateUserProfile(&newUpdateProfile, userPayload.UserID); err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to update user profile"))
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User profile updated successfully",
	})
}
