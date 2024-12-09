package repositories

import (
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfileRepository interface {
	CreateUserProfile(userProfile *models.UserProfile) error
	UpdateUserProfile(userProfileBody *models.UserProfile, userID uuid.UUID) error
}

type userProfileRepository struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) UserProfileRepository {
	return &userProfileRepository{db: db}
}

func (r *userProfileRepository) CreateUserProfile(userProfile *models.UserProfile) error {
	if err := r.db.Create(userProfile).Error; err != nil {
		return err
	}
	return nil
}

func (r *userProfileRepository) UpdateUserProfile(userProfileBody *models.UserProfile, userID uuid.UUID) error {
	var user models.UserProfile
	if err := r.db.Model(&user).Where("user_id = ?", userID).Updates(userProfileBody).Error; err != nil {
		return err
	}

	return nil
}
