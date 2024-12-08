package repositories

import (
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemTenorRepository interface {
	FindItemLimitByItemIDTenorID(itemID, tenorID uuid.UUID) (*models.ItemTenor, error)
}

type itemTenorRepository struct {
	db *gorm.DB
}

func NewItemLimitRepository(db *gorm.DB) ItemTenorRepository {
	return &itemTenorRepository{db: db}
}

func (r *itemTenorRepository) FindItemLimitByItemIDTenorID(itemID, tenorID uuid.UUID) (*models.ItemTenor, error) {
	var itemTenor models.ItemTenor
	if err := r.db.Where("item_id = ? AND tenor_id = ?", itemID, tenorID).First(&itemTenor).Error; err != nil {
		return nil, err
	}

	return &itemTenor, nil
}
