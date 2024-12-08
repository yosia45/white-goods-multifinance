package utils

import (
	"time"
	"white-goods-multifinace/constants"
)

func InvoiceGenerator(purchaseID string) string {
	return "INV-" + purchaseID + "/" + constants.RomanNumerals[int(time.Now().Month())] + "/" + time.Now().Format("2006")
}
