package dto

import "github.com/google/uuid"

type ItemResponse struct {
	ItemID      uuid.UUID   `json:"item_id"`
	Name        string      `json:"name"`
	NormalPrice float64     `json:"normal_price"`
	AdminFee    float64     `json:"admin_fee"`
	OTR         OTRResponse `json:"on_the_road"`
}
