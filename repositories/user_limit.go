package repositories

import (
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLimitRepository interface {
	CreateUserLimit(userLimits *[]models.UserLimit) error
	FindUserLimitByUserIDTenorID(userID, tenorID uuid.UUID) (*models.UserLimit, error)
	UpdateUserLimit(userLimit *models.UserLimit, userID uuid.UUID, tenorID uuid.UUID) error
	BulkUpdateUserLimit(userLimits *[]models.UserLimit, userID uuid.UUID) error
}

type userLimitRepository struct {
	db *gorm.DB
}

func NewUserLimitRepository(db *gorm.DB) UserLimitRepository {
	return &userLimitRepository{db: db}
}

func (r *userLimitRepository) CreateUserLimit(userLimits *[]models.UserLimit) error {
	if err := r.db.Create(userLimits).Error; err != nil {
		return err
	}
	return nil
}

func (r *userLimitRepository) FindUserLimitByUserIDTenorID(userID, tenorID uuid.UUID) (*models.UserLimit, error) {
	var userLimit models.UserLimit
	if err := r.db.Where("user_id = ? AND tenor_id = ?", userID, tenorID).First(&userLimit).Error; err != nil {
		return nil, err
	}

	return &userLimit, nil
}

func (r *userLimitRepository) BulkUpdateUserLimit(userLimits *[]models.UserLimit, userID uuid.UUID) error {
	if err := r.db.Model(&userLimits).Where("user_id = ?", userID).Updates(userLimits).Error; err != nil {
		return err
	}

	return nil
}

func (r *userLimitRepository) UpdateUserLimit(userLimit *models.UserLimit, userID uuid.UUID, tenorID uuid.UUID) error {
	var userLimitModel models.UserLimit
	if err := r.db.Model(&userLimitModel).Where("user_id = ? AND tenor_id = ?", userID, tenorID).Updates(userLimit).Error; err != nil {
		return err
	}

	return nil
}
