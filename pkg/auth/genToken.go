package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(username string) (string, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")
	jwtKey := []byte(JWT_SECRET)

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
