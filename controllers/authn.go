package controllers

import (
	"net/http"
	"white-goods-multifinace/dto"
	"white-goods-multifinace/middlewares"
	"white-goods-multifinace/models"
	"white-goods-multifinace/repositories"
	"white-goods-multifinace/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userRepo        repositories.UserRepository
	userProfileRepo repositories.UserProfileRepository
}

func NewUserController(userRepo repositories.UserRepository, userProfileRepo repositories.UserProfileRepository) *UserController {
	return &UserController{
		userRepo:        userRepo,
		userProfileRepo: userProfileRepo,
	}
}

func (uc *UserController) RegisterCustomer(c echo.Context) error {
	var customer dto.RegisterUserBody
	if err := c.Bind(&customer); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if customer.Email == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Email is required"))
	}

	if customer.Password == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Password is required"))
	}

	if customer.FullName == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Full name is required"))
	}

	if customer.NIK == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("NIK is required"))
	}

	if len(customer.NIK) != 16 {
		return utils.HandlerError(c, utils.NewBadRequestError("NIK must be 16 characters long"))
	}

	newCustomer := models.User{
		Email:    customer.Email,
		Password: customer.Password,
		Role:     "customer",
	}

	newCustomerProfile := models.UserProfile{
		FullName: customer.FullName,
		NIK:      customer.NIK,
	}

	if err := uc.userRepo.CreateUser(&newCustomer, &newCustomerProfile); err != nil {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Customer registered successfully"})
}

func (uc *UserController) Login(c echo.Context) error {
	var loginBody dto.LoginUserBody

	if err := c.Bind(&loginBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if loginBody.Email == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Email is required"))
	}

	if loginBody.Password == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Password is required"))
	}

	user, err := uc.userRepo.FindUserByEmail(loginBody.Email)
	if err != nil {
		return utils.HandlerError(c, utils.NewNotFoundError("User not found"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password))
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid email or password"))
	}

	token, err := middlewares.GenerateJWT(user.UserID, user.Role, 60)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
