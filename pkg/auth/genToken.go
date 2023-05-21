package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GenerateToken(username string, ID uint) (string, error) {
	JWT_SECRET := GetJWTSecret()
	jwtKey := []byte(JWT_SECRET)

	claims := jwt.MapClaims{
		"ID":       ID,
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
