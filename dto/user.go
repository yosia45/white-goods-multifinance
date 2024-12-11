package dto

import "github.com/google/uuid"

type RegisterUserBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type LoginUserBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserByEmailResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}

type UserByIDResponse struct {
	ID        string                    `json:"id"`
	Email     string                    `json:"email"`
	Role      string                    `json:"role"`
	Details   UserProfileResponse       `json:"details"`
	Limits    []UserLimitDetailResponse `json:"limits"`
	Purchases []GetAllUserPurchase      `json:"purchases"`
}
