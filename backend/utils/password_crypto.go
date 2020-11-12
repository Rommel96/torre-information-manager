package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidPassword(hashedPassword, planePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(planePassword))
}
