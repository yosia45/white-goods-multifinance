package models

type Tenor struct {
	BaseModel
	Duration   int         `json:"duration" gorm:"not null"`
	Interest   float64     `json:"interest" gorm:"not null"`
	IsDefault  bool        `json:"is_default" gorm:"not null"`
	ItemTenors []ItemTenor `json:"item_tenors" gorm:"foreignKey:TenorID"`
}
