package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashePassword returns the bcrypt hash of password

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Gagal untuk melakukan hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// Check Password if the provided password is correct or not
func checkPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}
