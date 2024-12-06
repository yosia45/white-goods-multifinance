package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Email       string      `json:"email" gorm:"unique;not null"`
	Password    string      `json:"password" gorm:"not null"`
	Role        string      `json:"role" gorm:"not null"`
	UserProfile UserProfile `json:"user_profile" gorm:"foreignKey:UserID"`
	UserLimits  []UserLimit `json:"user_limits" gorm:"foreignKey:UserID"`
	Purchases   []Purchase  `json:"purchases" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	u.Password = string(hashedPassword)
	return
}
