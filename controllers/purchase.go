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

type PurchaseController struct {
	userRepo      repositories.UserRepository
	userLimitRepo repositories.UserLimitRepository
	itemRepo      repositories.ItemRepository
	itemTenorRepo repositories.ItemTenorRepository
	purchaseRepo  repositories.PurchaseRepository
	tenorRepo     repositories.TenorRepository
}

func NewPurchaseController(userRepo repositories.UserRepository, userLimitRepo repositories.UserLimitRepository, itemRepo repositories.ItemRepository, itemTenorRepo repositories.ItemTenorRepository, purchaseRepo repositories.PurchaseRepository, tenorRepo repositories.TenorRepository) *PurchaseController {
	return &PurchaseController{
		userRepo:      userRepo,
		userLimitRepo: userLimitRepo,
		itemRepo:      itemRepo,
		itemTenorRepo: itemTenorRepo,
		purchaseRepo:  purchaseRepo,
		tenorRepo:     tenorRepo,
	}
}

func (pc *PurchaseController) CreatePurchase(c echo.Context) error {
	var purchaseBody dto.AddPurchaseBody
	userPayload := c.Get("userPayload").(*dto.JWTPayload)

	if err := c.Bind(&purchaseBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if purchaseBody.ItemID == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Item ID is required"))
	}

	if purchaseBody.TenorID == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Tenor ID is required"))
	}

	parsedItemID, err := uuid.Parse(purchaseBody.ItemID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid item ID"))
	}

	parsedTenorID, err := uuid.Parse(purchaseBody.TenorID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid tenor ID"))
	}

	userLimit, err := pc.userLimitRepo.FindUserLimitByUserIDTenorID(userPayload.UserID, parsedTenorID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("User or Tenor not found"))
	}

	itemTenor, err := pc.itemTenorRepo.FindItemLimitByItemIDTenorID(parsedItemID, parsedTenorID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Item or Tenor not found"))
	}

	foundItem, err := pc.itemRepo.FindItemByID(parsedItemID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Item not found"))
	}

	tenor, err := pc.tenorRepo.FindTenorByID(parsedTenorID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Tenor not found"))
	}

	if userLimit.CurrentBalance < foundItem.NormalPrice {
		return utils.HandlerError(c, utils.NewBadRequestError("Insufficient balance"))
	}

	monthlyInstallment := ((1+(itemTenor.Interest/100))*float64(foundItem.NormalPrice) + foundItem.AdminFee) / float64(tenor.Duration)

	newPurchase := models.Purchase{
		UserLimitID:    userLimit.ID,
		ItemTenorID:    itemTenor.ID,
		MonthlyPayment: monthlyInstallment,
		IsCompleted:    false,
	}

	if err := pc.purchaseRepo.CreatePurchase(&newPurchase); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Failed to create purchase"))
	}

	newUserCurrentBalance := userLimit.CurrentBalance - foundItem.NormalPrice

	newCurrentBalance := models.UserLimit{
		CurrentBalance: newUserCurrentBalance,
	}

	if err := pc.userLimitRepo.UpdateUserLimit(&newCurrentBalance, userPayload.UserID); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Failed to update user balance"))
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Purchase created successfully"})
}
