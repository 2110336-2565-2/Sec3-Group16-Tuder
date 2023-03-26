package middlewares

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
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
			ErrorHandler: func(c echo.Context, err error) error {
				return c.JSON(401, &schemas.SchemaErrorResponse{
					Success: false,
					Message: err.Error() + " : " + "Unauthorized",
					Error:   err,
				})
			},
		},
	)
}

func (a *authMiddleware) AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		if !ok {
			return c.JSON(401, &schemas.SchemaErrorResponse{
				Success: false,
				Message: "JWT token missing or invalid",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
		if !ok {
			return c.JSON(401, &schemas.SchemaErrorResponse{
				Success: false,
				Message: "failed to cast claims as jwt.MapClaims",
			})
		}
		if claims["role"] == "admin" {
			return next(c)
		}
		return c.JSON(401, &schemas.SchemaErrorResponse{
			Success: false,
			Message: "user unauthorized",
		})
	}
}

//
//func (m *authMiddleware) JWT(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		echojwt.
//	}
//}
