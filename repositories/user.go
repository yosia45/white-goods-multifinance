package repositories

import (
	"errors"
	"fmt"
	"white-goods-multifinace/dto"
	"white-goods-multifinace/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User, userProfile *models.UserProfile) error
	FindUserByEmail(email string) (*dto.UserByEmailResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User, userProfile *models.UserProfile) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Printf("Transaction failed: %v\n", r)
		}
	}()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	userProfile.UserID = user.ID

	if err := tx.Create(userProfile).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUserByEmail(email string) (*dto.UserByEmailResponse, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	response := dto.UserByEmailResponse{
		UserID:   user.ID,
		Password: user.Password,
		Role:     user.Role,
	}

	return &response, nil
}
