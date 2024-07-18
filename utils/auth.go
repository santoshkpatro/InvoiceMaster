package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// GenerateSalt generates a random salt of specified length
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword hashes a password with a given salt
func HashPassword(password, salt string) (string, error) {
	saltedPassword := password + salt
	hash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash checks if the provided password and salt match the hashed password
func CheckPasswordHash(password, salt, hash string) bool {
	saltedPassword := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedPassword))
	return err == nil
}
