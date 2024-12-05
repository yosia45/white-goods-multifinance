package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfile struct {
	BaseModel
	FullName       string    `json:"full_name" gorm:"not null"`
	LegalName      string    `json:"legal_name" gorm:"not null"`
	NIK            string    `json:"nik" gorm:"unique;not null;size:16"`
	BirthPlace     string    `json:"birth_place" gorm:"not null"`
	BirthDate      time.Time `json:"birth_date" gorm:"not null"`
	Salary         float64   `json:"salary" gorm:"not null"`
	KTPFilePath    string    `json:"ktp_file_path" gorm:"not null"`
	SelfieFilePath string    `json:"selfie_file_path" gorm:"not null"`
	UserID         uuid.UUID `json:"user_id" gorm:"not null;size:191"`
}

func (up *UserProfile) BeforeCreate(tx *gorm.DB) (err error) {
	up.ID = uuid.New()
	up.CreatedAt = time.Now()
	up.UpdatedAt = time.Now()

	return
}
