package dto

import "github.com/google/uuid"

type AddUserLimitBody struct {
	UserID      uuid.UUID `json:"user_id"`
	LimitTenor1 float64   `json:"limit_tenor_1"`
	LimitTenor2 float64   `json:"limit_tenor_2"`
	LimitTenor3 float64   `json:"limit_tenor_3"`
	LimitTenor6 float64   `json:"limit_tenor_6"`
}
