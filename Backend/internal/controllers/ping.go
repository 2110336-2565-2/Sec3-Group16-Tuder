package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func PingController(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
