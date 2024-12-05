package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Email                      string                      `json:"email" gorm:"unique;not null"`
	Password                   string                      `json:"password" gorm:"not null"`
	UserProfile                UserProfile                 `json:"user_profile" gorm:"foreignKey:UserID"`
	UserMoney                  UserMoney                   `json:"user_money" gorm:"foreignKey:UserID"`
	UserPurchasingInformations []UserPurchasingInformation `json:"user_purchasing_informations" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	u.Password = string(hashedPassword)
	return
}
