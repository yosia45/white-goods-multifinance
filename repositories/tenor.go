package repositories

import (
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TenorRepository interface {
	FindTenorByID(tenorID uuid.UUID) (*models.Tenor, error)
}

type tenorRepository struct {
	db *gorm.DB
}

func NewTenorRepository(db *gorm.DB) TenorRepository {
	return &tenorRepository{db: db}
}

func (r *tenorRepository) FindTenorByID(tenorID uuid.UUID) (*models.Tenor, error) {
	var tenor models.Tenor
	if err := r.db.Where("id = ?", tenorID).First(&tenor).Error; err != nil {
		return nil, err
	}

	return &tenor, nil
}
