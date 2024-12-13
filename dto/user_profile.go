package dto

import (
	"time"

	"github.com/google/uuid"
)

type UpdateUserProfileBody struct {
	FullName          string    `form:"full_name"`
	LegalName         string    `form:"legal_name"`
	BirthPlace        string    `form:"birth_place"`
	BirthDate         time.Time `form:"birth_date"`
	Salary            float64   `form:"salary"`
	KTPFilePathURL    string    `form:"ktp_file"`
	selfieFilePathURL string    `form:"selfie_file"`
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
