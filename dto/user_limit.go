package dto

import "github.com/google/uuid"

type AddUserLimitBody struct {
	UserID         string  `json:"user_id"`
	TenorID        string  `json:"tenor_id"`
	Limit          float64 `json:"limit"`
	CurrentBalance float64 `json:"current_balance"`
}

type UserLimitBody struct {
	UserID         uuid.UUID `json:"user_id"`
	TenorID        uuid.UUID `json:"tenor_id"`
	Limit          float64   `json:"limit"`
	CurrentBalance float64   `json:"current_balance"`
}

type UpdateCurrentBalanceUserLimit struct {
	UserLimitID    uuid.UUID `json:"user_limit_id"`
	CurrentBalance float64   `json:"current_balance"`
}

type UserLimitDetailResponse struct {
	Duration uint    `json:"duration"`
	Limit    float64 `json:"limit"`
}
