package dto

import "github.com/google/uuid"

type TenorResponse struct {
	TenorID  uuid.UUID `json:"tenor_id"`
	Duration int       `json:"duration"`
}
