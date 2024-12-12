package repositories

import (
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	FindTransactionByPurchaseID(purchaseID uuid.UUID) (*[]models.Transaction, error)
	Begin() *gorm.DB
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
