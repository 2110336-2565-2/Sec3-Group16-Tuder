package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string, isExpire bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["iat"] = time.Now().Unix()
	if isExpire {
		claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateLoginToken(username string, role string, isExpire bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	if isExpire {
		claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
