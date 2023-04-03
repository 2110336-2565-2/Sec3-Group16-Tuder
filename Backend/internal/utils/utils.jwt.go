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

func GenerateEmailToken(userid uuid.UUID, isExpire bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = userid
	claims["iat"] = time.Now().Unix()
	if isExpire {
		exp, _ := s.Atoi(os.Getenv("JWT_EXPIRES_MINUTES_EMAIL"))
		claims["exp"] = time.Now().Add(time.Duration(exp) * time.Minute).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// Decode the token and return the userid
func DecodeToken(token string) (uuid.UUID, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	userid, err := uuid.Parse(claims["userid"].(string))
	if err != nil {
		return uuid.Nil, err
	}
	return userid, nil
}
