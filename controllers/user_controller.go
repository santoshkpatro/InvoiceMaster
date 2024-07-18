package controllers

import (
	"InvoiceMaster/config"
	"InvoiceMaster/models"
	"InvoiceMaster/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func RegisterUser(c echo.Context) error {
	registerUser := new(models.RegisterUserModel)
	if err := c.Bind(registerUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := validate.Struct(registerUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Check if email already exists
	var existingUser models.User
	err := config.DB.Get(&existingUser, "SELECT id FROM users WHERE email = $1", registerUser.Email)
	if err == nil {
		// User with this email already exists
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email already registered"})
	}

	salt, _ := utils.GenerateSalt(16)
	password_hash, _ := utils.HashPassword(registerUser.Password, salt)

	tx := config.DB.MustBegin()
	tx.MustExec("INSERT INTO users (full_name, email, salt, password_hash) VALUES ($1, $2, $3, $4);", registerUser.FullName, registerUser.Email, salt, password_hash)
	tx.Commit()

	return c.String(http.StatusCreated, "Register Success")
}
