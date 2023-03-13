package controllers

import (
	"net/http"

	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	services "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type ControllerChangePassword interface {
	ChangePassword(c echo.Context) error
}

type controllerChangePassword struct {
	service services.ServiceChangePassword
}

func NewControllerChangePassword(s services.ServiceChangePassword) ControllerChangePassword {
	return &controllerChangePassword{
		service: s,
	}
}

func (cC *controllerChangePassword) ChangePassword(c echo.Context) error {
	user := &schemas.SchemaChangePassword{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	err := cC.service.ChangePassword(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, "success")

}
