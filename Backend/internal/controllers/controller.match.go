package controllers

import (
	"fmt"
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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
	claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
	studentId, ok := claims["studentid"]
	if !ok {
		return fmt.Errorf("claim \"studentid\" does not exist")
	}
	uR.Student_id, err = uuid.Parse(studentId.(string))
	if err != nil {
		return fmt.Errorf("studentId from jwt token is invalid")
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
