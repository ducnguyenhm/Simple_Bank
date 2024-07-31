package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bscrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)

	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, HashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))


}
