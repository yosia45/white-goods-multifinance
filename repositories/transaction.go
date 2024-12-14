package repositories

import (
	"fmt"
	"sync"
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	FindTransactionByPurchaseID(purchaseID uuid.UUID) (*[]models.Transaction, error)
	CreateTransactionCouncurrentTransaction(transaction *models.Transaction, userID uuid.UUID, tenorID uuid.UUID, itemNormalPrice float64) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(transaction *models.Transaction) error {
	if err := r.db.Create(transaction).Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) CreateTransactionCouncurrentTransaction(transaction *models.Transaction, userID uuid.UUID, tenorID uuid.UUID, itemNormalPrice float64) error {
	var wg sync.WaitGroup
	var mu sync.Mutex

	errCh := make(chan error, 1)

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback jika terjadi panic
			errCh <- fmt.Errorf("transaction failed: %v", r)
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		tx := r.db.Begin()
		if tx.Error != nil {
			errCh <- tx.Error
			return
		}

		if err := tx.Create(transaction).Error; err != nil {
			tx.Rollback()
			errCh <- err
			return
		}

		if err := tx.Commit().Error; err != nil {
			errCh <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		tx := r.db.Begin()
		if tx.Error != nil {
			errCh <- tx.Error
			return
		}

		var userLimit models.UserLimit
		if err := tx.Where("user_id = ? AND tenor_id = ?", userID, tenorID).First(&userLimit).Error; err != nil {
			errCh <- err
			return
		}

		newCurrentBalance := userLimit.CurrentBalance + itemNormalPrice

		updatedUserLimit := models.UserLimit{
			CurrentBalance: newCurrentBalance,
		}

		mu.Lock()
		defer mu.Unlock()

		if err := tx.Model(&userLimit).Where("user_id = ? AND tenor_id = ?", userID, tenorID).Updates(&updatedUserLimit).Error; err != nil {
			tx.Rollback()
			errCh <- err
			return
		}

		if err := tx.Commit().Error; err != nil {
			errCh <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		tx := r.db.Begin()
		if tx.Error != nil {
			errCh <- tx.Error
			return
		}

		updatedPurchase := models.Purchase{
			IsCompleted: true,
		}

		if err := tx.Model(models.Purchase{}).Where("id = ?", transaction.PurchaseID).Updates(&updatedPurchase).Error; err != nil {
			tx.Rollback()
			errCh <- err
			return
		}

		if err := tx.Commit().Error; err != nil {
			errCh <- err
		}
	}()

	wg.Wait()

	close(errCh)

	for err := range errCh {
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) FindTransactionByPurchaseID(purchaseID uuid.UUID) (*[]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Where("purchase_id = ?", purchaseID).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (r *transactionRepository) Begin() *gorm.DB {
	return r.db.Begin()
}
