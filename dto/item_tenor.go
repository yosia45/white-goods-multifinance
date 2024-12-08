package dto

import "github.com/google/uuid"

type ItemTenorResponse struct {
	ItemTenorID  uuid.UUID             `json:"item_tenor_id"`
	Interest     float64               `json:"interest"`
	Item         ItemResponse          `json:"item"`
	Tenor        TenorResponse         `json:"tenor"`
	Transactions []TransactionResponse `json:"transactions"`
}
