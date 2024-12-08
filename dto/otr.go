package dto

import "github.com/google/uuid"

type OTRResponse struct {
	OTRID uuid.UUID `json:"otr_id"`
	Name  string    `json:"name"`
}
