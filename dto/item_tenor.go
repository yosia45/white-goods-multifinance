package dto

import "github.com/google/uuid"

type AddItemTenorBody struct {
	ItemID         uuid.UUID `json:"item_id"`
	AmountTenor1   float64   `json:"amount_tenor_1"`
	AmountTenor2   float64   `json:"amount_tenor_2"`
	AmountTenor3   float64   `json:"amount_tenor_3"`
	AmountTenor6   float64   `json:"amount_tenor_6"`
	InterestTenor1 float64   `json:"interest_tenor_1"`
	InterestTenor2 float64   `json:"interest_tenor_2"`
	InterestTenor3 float64   `json:"interest_tenor_3"`
	InterestTenor6 float64   `json:"interest_tenor_6"`
}
