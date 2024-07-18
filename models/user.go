package models

type User struct {
	ID           int    `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Salt         string `json:"salt"`
	PasswordHash string `json:"password_hash"`
}

type RegisterUserModel struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
