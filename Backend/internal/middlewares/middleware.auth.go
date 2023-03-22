package middlewares

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
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

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role")
		if role == "admin" {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}

//
//func (m *authMiddleware) JWT(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		echojwt.
//	}
//}
