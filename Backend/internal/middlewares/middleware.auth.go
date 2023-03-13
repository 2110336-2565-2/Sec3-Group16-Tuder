package middlewares

import (
	"github.com/labstack/echo/v4"
)

import (
	"github.com/labstack/echo-jwt/v4"
)

type authMiddleware struct {
	secret string
}

func NewAuthMiddleware(s string) *authMiddleware {
	return &authMiddleware{secret: s}
}

func (a *authMiddleware) JWT() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(a.secret),
		},
	)
}

//
//func (m *authMiddleware) JWT(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		echojwt.
//	}
//}
