package repositories

import (
	"errors"
	"white-goods-multifinace/dto"
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (uuid.UUID, error)
	FindUserByEmail(email string) (*dto.UserByEmailResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) (uuid.UUID, error) {
	if err := r.db.Create(user).Error; err != nil {
		return uuid.Nil, err
	}
	return user.ID, nil
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
