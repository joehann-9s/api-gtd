package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

func GetIDFromClaims(r *http.Request) uint {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	userID := uint(claims["ID"].(float64))
	return userID
}
