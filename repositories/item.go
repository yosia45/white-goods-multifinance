package repositories

import (
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindItemByID(itemID uuid.UUID) (*models.Item, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) FindItemByID(itemID uuid.UUID) (*models.Item, error) {
	var item models.Item
	if err := r.db.Where("id = ?", itemID).First(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}
