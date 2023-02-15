package controllers

import (
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type controllerResetPassword struct {
	service service.ServiceResetPassword
}

func NewControllerResetPassword(s service.ServiceResetPassword) *controllerResetPassword {
	return &controllerResetPassword{service: s}
}

func (cR *controllerResetPassword) ResetPassword(c echo.Context) (err error) {
	return nil
}
