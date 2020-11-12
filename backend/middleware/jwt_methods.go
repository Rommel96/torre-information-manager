package middleware

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(email interface{}) (string, error) {
	claims := jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 168).Unix(),
		"iat":   time.Now().Unix(),
		"email": email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t, err
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected token signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
