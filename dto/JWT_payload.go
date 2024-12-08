package dto

import "github.com/google/uuid"

type JWTPayload struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}
