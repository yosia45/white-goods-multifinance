package dto

import "github.com/google/uuid"

type AddPurchaseBody struct {
	TenorID string `json:"tenor_id"`
	ItemID  string `json:"item_id"`
}

type PurchaseBody struct {
	UserID         uuid.UUID `json:"user_id"`
	ItemTenorID    uuid.UUID `json:"item_tenor_id"`
	MonthlyPayment float64   `json:"monthly_payment"`
	IsCompleted    bool      `json:"is_completed"`
}

type GetAllUserPurchase struct {
	PurchaseID     string                  `json:"purchase_id"`
	MonthlyPayment float64                 `json:"monthly_payment"`
	IsCompleted    bool                    `json:"is_completed"`
	Purchases      GetAllPurchaseItemTenor `json:"purchases"`
}

type PurchaseByIDResponse struct {
	PurchaseID     uuid.UUID         `json:"purchase_id"`
	IsCompleted    bool              `json:"is_completed"`
	MonthlyPayment float64           `json:"monthly_payment"`
	ItemTenor      ItemTenorResponse `json:"item_tenor"`
}
