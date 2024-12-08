package dto

import (
	"time"

	"github.com/google/uuid"
)

type AddTransactionBody struct {
	PurchaseID  string    `json:"purchase_id"`
	TotalAmount float64   `json:"total_amount"`
	PaymentDate time.Time `json:"payment_date"`
}

type TransactionBody struct {
	PurchaseID    uuid.UUID `json:"purchase_id"`
	TotalAmount   float64   `json:"total_amount"`
	PaymentDate   time.Time `json:"payment_date"`
	InvoiceNumber string    `json:"invoice_number"`
}

type TransactionResponse struct {
	TransactionID string    `json:"transaction_id"`
	TotalAmount   float64   `json:"total_amount"`
	PaymentDate   time.Time `json:"payment_date"`
	InvoiceNumber string    `json:"invoice_number"`
}
