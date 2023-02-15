package controllers

import (
	"fmt"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type controllerRegister struct {
	service service.ServiceRegister
}

func NewControllerRegister(s service.ServiceRegister) *controllerRegister {
	return &controllerRegister{service: s}
}

func (cR *controllerRegister) RegisterUser(c echo.Context) (err error) {
	uR := &schema.SchemaRegister{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaRegisterResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}

	if err = cR.service.RegisterService(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaRegisterResponse{
			Success: false,
			Message: fmt.Sprint(err),
			Error:   err,
		})
		return
	}

	c.JSON(http.StatusOK, schema.SchemaRegisterResponse{
		Success: true,
		Message: "Register successfully",
	})
	return nil
}
