package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(username string, isExpire bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	if isExpire {
		claims["exp"] = time.Now().Add(time.Minute * 10)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}