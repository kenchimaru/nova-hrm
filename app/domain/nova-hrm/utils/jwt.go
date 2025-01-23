package utils

import (
	"nova-hrm/app/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username string, role string) string {
	secretKey := []byte(config.GetEnv("JWT_SECRET", "secret"))

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hour
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return tokenString
}
