package dto

import "time"

type UpdateUserProfileBody struct {
	FullName       string    `json:"full_name"`
	LegalName      string    `json:"legal_name"`
	NIK            string    `json:"nik"`
	BirthPlace     string    `json:"birth_place"`
	BirthDate      time.Time `json:"birth_date"`
	Salary         float64   `json:"salary"`
	KTPFilePath    string    `json:"ktp_file_path"`
	SelfieFilePath string    `json:"selfie_file_path"`
}
