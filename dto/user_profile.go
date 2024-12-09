package dto

import (
	"time"

	"github.com/google/uuid"
)

type UpdateUserProfileBody struct {
	FullName       string    `json:"full_name"`
	LegalName      string    `json:"legal_name"`
	NIK            string    `json:"nik"`
	BirthPlace     string    `json:"birth_place"`
	BirthDate      time.Time `json:"birth_date"`
	Salary         float64   `json:"salary"`
	KTPFilePath    string    `json:"ktp_file_path"`
	SelfieFilePath string    `json:"selfie_file_path"`
}

type AddUserProfileBody struct {
	FullName string    `json:"full_name"`
	UserID   uuid.UUID `json:"user_id"`
}

type PurchaseUserResponse struct {
	UserID    uuid.UUID `json:"user_id"`
	FullName  string    `json:"full_name"`
	LegalName string    `json:"legal_name"`
	NIK       string    `json:"nik"`
	Salary    float64   `json:"salary"`
}
type UserProfileResponse struct {
	FullName   string    `json:"full_name"`
	LegalName  string    `json:"legal_name"`
	NIK        string    `json:"nik"`
	BirthPlace string    `json:"birth_place"`
	BirthDate  time.Time `json:"birth_date"`
	Salary     float64   `json:"salary"`
}
