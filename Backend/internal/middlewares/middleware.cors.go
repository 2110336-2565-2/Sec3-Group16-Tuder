package middlewares

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func CorsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request().Method == "OPTIONS" {
			return c.NoContent(http.StatusOK)
		}
		return next(c)
	}
}
