package dto

type RegisterUserBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type LoginUserBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
