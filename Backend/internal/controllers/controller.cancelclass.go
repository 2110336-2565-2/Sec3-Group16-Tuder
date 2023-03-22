package controllers

import (
	"net/http"

	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	services "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type ControllerCancelClass interface {
	CancelClass(c echo.Context) error
}

type controllerCancelClass struct {
	service services.ServiceCancelClass
}

func NewControllerCancelClass(s services.ServiceCancelClass) ControllerCancelClass {
	return &controllerCancelClass{
		service: s,
	}
}

func (cC *controllerCancelClass) CancelClass(c echo.Context) error {
	request := &schemas.SchemaCancelClass{}
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.SchemaResponses{
			Success: false,
			Message: "invalid request",
			Data:    err.Error(),
		})

		return err
	}

	err := cC.service.CancelClass(request)
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
		Message: "Cancel request has already been sent to admin",
		Data:    nil,
	})
	return nil

}
