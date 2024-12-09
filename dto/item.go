package dto

import "github.com/google/uuid"

type AddItemBody struct {
	Name           string  `json:"name"`
	NormalPrice    float64 `json:"normal_price"`
	AdminFee       float64 `json:"admin_fee"`
	OTRID          uint    `json:"otr_id"`
	Tenor1ID       string  `json:"tenor_1_id"`
	Tenor2ID       string  `json:"tenor_2_id"`
	Tenor3ID       string  `json:"tenor_3_id"`
	Tenor6ID       string  `json:"tenor_6_id"`
	InterestTenor1 float64 `json:"interest_tenor_1"`
	InterestTenor2 float64 `json:"interest_tenor_2"`
	InterestTenor3 float64 `json:"interest_tenor_3"`
	InterestTenor6 float64 `json:"interest_tenor_6"`
}

type ItemResponse struct {
	ItemID      uuid.UUID   `json:"item_id"`
	Name        string      `json:"name"`
	NormalPrice float64     `json:"normal_price"`
	AdminFee    float64     `json:"admin_fee"`
	OTR         OTRResponse `json:"on_the_road"`
}
