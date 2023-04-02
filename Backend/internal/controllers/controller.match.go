package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type controllerMatch struct {
	service service.ServiceMatch
}

func NewControllerMatch(service service.ServiceMatch) *controllerMatch {
	return &controllerMatch{service: service}
}

func (cR *controllerMatch) CreateMatch(c echo.Context) (err error) {
	uR := &schema.SchemaCreateMatch{}
	if err := c.Bind(&uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return err
	}
	match, err := cR.service.CreateMatch(uR)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return err
	}
	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Create match successfully",
		Data:    match,
	})
	return
}
