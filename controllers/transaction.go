package controllers

import (
	"fmt"
	"net/http"
	"sync"
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
	userPayload := c.Get("userPayload").(*dto.JWTPayload)
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

	invoiceNumber := utils.InvoiceGenerator(transactionBody.PurchaseID)

	newTransaction := models.Transaction{
		PurchaseID:    parsedPurchaseID,
		TotalAmount:   transactionBody.TotalAmount,
		PaymentDate:   transactionBody.PaymentDate,
		InvoiceNumber: invoiceNumber,
	}

	if (purchase.ItemTenor.Tenor.Duration - 1) == len(purchase.ItemTenor.Transactions) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		errCh := make(chan error, 3)

		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := tc.transactionRepo.CreateTransaction(&newTransaction); err != nil {
				errCh <- fmt.Errorf(err.Error())
				return
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			userLimit, err := tc.userLimitRepo.FindUserLimitByUserIDTenorID(userPayload.UserID, purchase.ItemTenor.Tenor.TenorID)
			if err != nil {
				errCh <- fmt.Errorf(err.Error())
				return
			}

			newCurrentBalance := userLimit.CurrentBalance + purchase.ItemTenor.Item.NormalPrice

			updatedUserLimit := models.UserLimit{
				CurrentBalance: newCurrentBalance,
			}

			mu.Lock()
			defer mu.Unlock()
			if err := tc.userLimitRepo.UpdateUserLimit(&updatedUserLimit, userPayload.UserID, purchase.ItemTenor.Tenor.TenorID); err != nil {
				errCh <- fmt.Errorf(err.Error())
				return
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			updatePurchase := models.Purchase{
				IsCompleted: true,
			}

			if err := tc.purchaseRepo.UpdatePurchaseByID(&updatePurchase, parsedPurchaseID); err != nil {
				errCh <- fmt.Errorf(err.Error())
				return
			}
		}()

		wg.Wait()

		close(errCh)

		for err := range errCh {
			if err != nil {
				return utils.HandlerError(c, utils.NewInternalError(err.Error()))
			}
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"message": "Transaction created successfully and purchased is completed",
		})
	} else if purchase.ItemTenor.Tenor.Duration > len(purchase.ItemTenor.Transactions) {
		if err := tc.transactionRepo.CreateTransaction(&newTransaction); err != nil {
			return utils.HandlerError(c, utils.NewInternalError(err.Error()))
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"message": "Transaction created successfully",
		})
	} else {
		return utils.HandlerError(c, utils.NewBadRequestError("Transaction is already completed"))
	}
}
