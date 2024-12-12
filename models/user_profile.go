package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfile struct {
	BaseModel
	FullName       string     `json:"full_name" gorm:"not null"`
	LegalName      *string    `json:"legal_name"`
	NIK            string     `json:"nik" gorm:"unique;size:16"`
	BirthPlace     *string    `json:"birth_place"`
	BirthDate      *time.Time `json:"birth_date"`
	Salary         *float64   `json:"salary"`
	KTPFilePath    *string    `json:"ktp_file_path"`
	SelfieFilePath *string    `json:"selfie_file_path"`
	UserID         uuid.UUID  `json:"user_id" gorm:"not null;size:191"`
}

func (up *UserProfile) BeforeCreate(tx *gorm.DB) (err error) {
	up.ID = uuid.New()
	up.CreatedAt = time.Now()
	up.UpdatedAt = time.Now()

	return
}
