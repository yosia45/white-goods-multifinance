package repositories

import (
	"white-goods-multifinace/dto"
	"white-goods-multifinace/models"

	"gorm.io/gorm"
)

type UserProfileRepository interface {
	CreateUserProfile(userProfile *models.UserProfile) error
	UpdateUserProfile(userProfileBody *dto.UpdateUserProfileBody) error
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

func (r *userProfileRepository) UpdateUserProfile(userProfileBody *dto.UpdateUserProfileBody) error {
	var user models.UserProfile
	if err := r.db.Model(&user).Updates(userProfileBody).Error; err != nil {
		return err
	}
	return nil
}
