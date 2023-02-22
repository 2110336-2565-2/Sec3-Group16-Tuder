package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type controllerTutor struct {
	service service.ServiceTutor
}

func NewControllerTutor(service service.ServiceTutor) *controllerTutor {
	return &controllerTutor{service: service}
}

func (cR *controllerTutor) GetTutorByUsername(c echo.Context) (err error) {
	uR := &schema.SchemaGetTutor{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}
	tutor, err := cR.service.GetTutorByUsername(uR)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return
	}
	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Get tutor successfully",
		Data:    tutor,
	})
	return nil
}

func (cR *controllerTutor) GetTutors(c echo.Context) (err error) {
	tutors, err := cR.service.GetTutors()
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return
	}
	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Get tutors successfully",
		Data:    tutors,
	})
	return nil
}

func (cR *controllerTutor) CreateTutor(c echo.Context) (err error) {
	uR := &schema.SchemaCreateTutor{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}

	tutor, err := cR.service.CreateTutor(uR)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return
	}

	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Create tutor successfully",
		Data:    tutor,
	})
	return nil
}

func (cR *controllerTutor) UpdateTutor(c echo.Context) (err error) {
	uR := &schema.SchemaUpdateTutor{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}

	tutor, err := cR.service.UpdateTutor(uR)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return
	}

	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Update tutor successfully",
		Data:    tutor,
	})
	return nil
}

func (cR *controllerTutor) DeleteTutor(c echo.Context) (err error) {
	uR := &schema.SchemaDeleteTutor{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}

	err = cR.service.DeleteTutor(uR)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return
	}

	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Delete tutor successfully",
		Data:    nil,
	})
	return nil
}
