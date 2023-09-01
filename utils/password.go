package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	Hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not has password %w", err)
	}

	return string(Hashedpassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
