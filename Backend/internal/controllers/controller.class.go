package controllers

import (
	"net/http"

	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	services "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type ControllerClass interface {
	GetCancellingClasses(echo.Context) error
	CancelClass(c echo.Context) error
	ApproveClassCancellation(c echo.Context) error
}

type controllerClass struct {
	service services.ServiceClass
}

func NewControllerClass(s services.ServiceClass) ControllerClass {
	return &controllerClass{
		service: s,
	}
}

func (cC *controllerClass) GetCancellingClasses(c echo.Context) error {
	classes, err := cC.service.GetCancellingClasses()
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return err
	}
	c.JSON(http.StatusOK, schemas.SchemaResponses{
		Success: true,
		Message: "Get classes successfully",
		Data:    classes,
	})
	return nil
}

func (cC *controllerClass) CancelClass(c echo.Context) error {
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

func (cC *controllerClass) ApproveClassCancellation(c echo.Context) error {

	cancelSchema := &schemas.SchemaCancelClass{}
	if err := c.Bind(cancelSchema); err != nil {
		c.JSON(http.StatusBadRequest, schemas.SchemaResponses{
			Success: false,
			Message: "invalid request",
			Data:    err.Error(),
		})
	}

	err := cC.service.ApproveClassCancellation(cancelSchema)
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
		Message: "Class cancellation has been approved",
		Data:    nil,
	})
	return nil

}
