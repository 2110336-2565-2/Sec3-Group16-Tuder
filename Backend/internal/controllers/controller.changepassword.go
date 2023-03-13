package controllers

import (
	"net/http"

	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	services "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type ControllerChangePassword interface {
	ChangePassword(c echo.Context) error
	CheckPassword(c echo.Context) error
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
		c.JSON(http.StatusBadRequest, schemas.SchemaResponses{
			Success: false,
			Message: "invalid request",
			Data:    err.Error(),
		})

		return err
	}
	err := cC.service.ChangePassword(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.SchemaResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schemas.SchemaResponses{
		Success: true,
		Message: "Change password successfully",
		Data:    nil,
	})
	return nil

}

func (cC *controllerChangePassword) CheckPassword(c echo.Context) error {
	user := &schemas.SchemaCheckPassword{}
	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, schemas.SchemaResponses{
			Success: false,
			Message: "invalid request",
			Data:    err.Error(),
		})
		return err
	}
	err := cC.service.CheckPassword(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.SchemaResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schemas.SchemaResponses{
		Success: true,
		Message: "Correct password",
		Data:    nil,
	})
	return nil

}
