package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenor struct {
	BaseModel
	Duration      int           `json:"duration" gorm:"not null"`
	ItemTenors    []ItemTenor   `json:"item_tenors" gorm:"foreignKey:TenorID"`
	UserLimits    []UserLimit   `json:"user_limits" gorm:"foreignKey:TenorID"`
	Transacations []Transaction `json:"transactions" gorm:"foreignKey:TenorID"`
}

func (t *Tenor) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	return
}
