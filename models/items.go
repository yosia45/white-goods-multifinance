package models

type Item struct {
	BaseModel
	Name        string      `json:"name" gorm:"not null"`
	NormalPrice float64     `json:"normal_price" gorm:"not null"`
	ItemTenors  []ItemTenor `json:"item_tenors" gorm:"foreignKey:ItemID"`
}
