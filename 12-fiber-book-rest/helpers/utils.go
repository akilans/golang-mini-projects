package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Log erros
func LogError(err error) {
	log.Printf("Error - %v \n", err.Error())
}

// Generate Hash Password
func GenerateHashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	return string(bytes), err
}

// Compare password
func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
