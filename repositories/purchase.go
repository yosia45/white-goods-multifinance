package repositories

import (
	"white-goods-multifinace/dto"
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchaseRepository interface {
	CreatePurchase(purchase *models.Purchase) error
	FindPurchaseByID(purchaseID uuid.UUID) (*dto.PurchaseByIDResponse, error)
	FindAllPurchase(userID uuid.UUID) (*[]dto.UserPurchaseResponse, error)
}

type purchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &purchaseRepository{db: db}
}

func (r *purchaseRepository) CreatePurchase(purchase *models.Purchase) error {
	if err := r.db.Create(purchase).Error; err != nil {
		return err
	}

	return nil
}

func (r *purchaseRepository) FindAllPurchase(userID uuid.UUID) (*[]dto.UserPurchaseResponse, error) {
	var purchases []models.Purchase
	// if err := r.db.Table("purchases p").
	// 	Select("p.id, p.status, p.monthly_payment, it.id, it.amount, it.interest, i.id, i.name, o.id, o.name, tn.id, tn.duration, t.id, t.total_amount, t.payment_date, t.invoice_number, t.status").
	// 	Joins("JOIN transactions t ON p.id = t.purchase_id").
	// 	Joins("JOIN item_tenors it ON p.item_tenor_id = it.id").
	// 	Joins("JOIN items i ON it.item_id = i.id").
	// 	Joins("JOIN tenors tn ON it.tenor_id = tn.id").
	// 	Joins("JOIN otrs o ON i.otr_id = o.id").
	// 	Where("p.user_id = ?", userID).
	// 	Find(&purchases).Error; err != nil {
	// 	return nil, err
	// }

	if err := r.db.Table("purchases p").
		Select("p.id, p.monthly_payment, p.is_completed").
		Joins("JOIN user_limits ul ON p.user_limit_id = ul.id").
		Joins("JOIN transactions t ON p.id = t.purchase_id").
		Where("ul.user_id = ?", userID).
		Find(&purchases).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *purchaseRepository) FindPurchaseByID(purchaseID uuid.UUID) (*dto.PurchaseByIDResponse, error) {
	var purchase models.Purchase
	if err := r.db.Table("purchases p").
		Select("p.id, p.monthly_payment, p.is_completed").
		Joins("JOIN user_limits ul ON p.user_limit_id = ul.id").
		Joins("JOIN transactions t ON p.id = t.purchase_id").
		Where("id = ?", purchaseID).
		Find(&purchase).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
