package repositories

import (
	"errors"
	"white-goods-multifinace/dto"
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfileRepository interface {
	CreateUserProfile(userProfile *models.UserProfile) error
	// FindUserProfileByNIK(nik string) (*dto.UserProfileResponse, error)
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

func (r *userProfileRepository) FindUserProfileByNIK(nik string) (*dto.UserProfileResponse, error) {
	var userProfileModel models.UserProfile
	if err := r.db.Where("nik = ?", nik).First(&userProfileModel).Error; err != nil {
		return nil, errors.New("user profile not found")
	}

	response := dto.UserProfileResponse{
		FullName:   userProfileModel.FullName,
		LegalName:  *userProfileModel.LegalName,
		NIK:        userProfileModel.NIK,
		BirthPlace: *userProfileModel.BirthPlace,
		BirthDate:  *userProfileModel.BirthDate,
		Salary:     *userProfileModel.Salary,
	}

	return &response, nil
}

func (r *userProfileRepository) UpdateUserProfile(userProfileBody *models.UserProfile, userID uuid.UUID) error {
	var user models.UserProfile
	if err := r.db.Model(&user).Where("user_id = ?", userID).Updates(userProfileBody).Error; err != nil {
		return err
	}

	return nil
}
