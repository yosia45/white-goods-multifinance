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

type TransactionController struct {
	transactionRepo repositories.TransactionRepository
	purchaseRepo    repositories.PurchaseRepository
	itemRepo        repositories.ItemRepository
	userLimitRepo   repositories.UserLimitRepository
}

func NewTransactionController(transactionRepo repositories.TransactionRepository, purchaseRepo repositories.PurchaseRepository, itemRepo repositories.ItemRepository, userLimitRepo repositories.UserLimitRepository) *TransactionController {
	return &TransactionController{
		transactionRepo: transactionRepo,
		purchaseRepo:    purchaseRepo,
		itemRepo:        itemRepo,
		userLimitRepo:   userLimitRepo,
	}
}

func (tc *TransactionController) CreateTransaction(c echo.Context) error {
	var transactionBody dto.AddTransactionBody
	userPayload := c.Get("userPayload").(dto.JWTPayload)
	if err := c.Bind(&transactionBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if transactionBody.PurchaseID == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Purchase ID is required"))
	}

	if transactionBody.PaymentDate.IsZero() {
		return utils.HandlerError(c, utils.NewBadRequestError("Payment date is required"))
	}

	if transactionBody.TotalAmount == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Total amount is required"))
	}

	parsedPurchaseID, err := uuid.Parse(transactionBody.PurchaseID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid purchase ID"))
	}

	purchase, err := tc.purchaseRepo.FindPurchaseByID(parsedPurchaseID)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	if purchase.MonthlyPayment != transactionBody.TotalAmount {
		return utils.HandlerError(c, utils.NewBadRequestError("Total amount must be equal to monthly payment"))
	}

	if purchase.ItemTenor.Tenor.Duration > len(purchase.ItemTenor.Transactions) {
		return utils.HandlerError(c, utils.NewBadRequestError("Transaction is already completed"))
	}

	invoiceNumber := utils.InvoiceGenerator(transactionBody.PurchaseID)

	newTransaction := models.Transaction{
		PurchaseID:    parsedPurchaseID,
		TotalAmount:   transactionBody.TotalAmount,
		PaymentDate:   transactionBody.PaymentDate,
		InvoiceNumber: invoiceNumber,
	}

	if err := tc.transactionRepo.CreateTransaction(newTransaction); err != nil {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	transactions, err := tc.transactionRepo.FindTransactionByPurchaseID(parsedPurchaseID)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError(err.Error()))
	}

	if len(*transactions) == purchase.ItemTenor.Tenor.Duration {
		userLimit, err := tc.userLimitRepo.FindUserLimitByUserIDTenorID(userPayload.UserID, purchase.ItemTenor.Tenor.TenorID)
		if err != nil {
			return utils.HandlerError(c, utils.NewInternalError(err.Error()))
		}

		newCurrentBalance := userLimit.CurrentBalance + purchase.ItemTenor.Item.NormalPrice

		updatedUserLimit := models.UserLimit{
			CurrentBalance: newCurrentBalance,
		}

		if err := tc.userLimitRepo.UpdateUserLimit(&updatedUserLimit, userPayload.UserID); err != nil {
			return utils.HandlerError(c, utils.NewInternalError(err.Error()))
		}
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Transaction created successfully",
	})
}
