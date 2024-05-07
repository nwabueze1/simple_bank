package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Cost factor: adjust based on desired security and performance trade-off
	cost := bcrypt.DefaultCost

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
			return "", err
	}

	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
			return errors.New("password does not match")
	}

	return nil
}

func IsPasswordMatch(hashed, password string)bool{
	err := verifyPassword(hashed, password)

	return err == nil
}
