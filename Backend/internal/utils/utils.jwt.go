package utils

import (
	"os"
	s "strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateLoginToken(customClaims map[string]interface{}, isExpire bool) (string, error) {
	claims := jwt.MapClaims{}
	// JWT required claims
	claims["authorized"] = true
	claims["iat"] = time.Now().Unix()
	if isExpire {
		exp, _ := s.Atoi(os.Getenv("JWT_EXPIRES_MINUTES"))
		claims["exp"] = time.Now().Add(time.Duration(exp) * time.Minute).Unix()
	}

	// add claims
	for k, v := range customClaims {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
