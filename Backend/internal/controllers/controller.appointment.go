package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type controllerAppointment struct {
	service service.ServiceAppointment
}

func NewControllerAppointment(service service.ServiceAppointment) *controllerAppointment {
	return &controllerAppointment{service: service}
}

func (cR *controllerAppointment) GetAppointmentByStudentID(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("student_id"))
	uR := &schema.SchemaGetAppointmentByStudentID{
		StudentID: id,
	}

	if err := c.Bind(&uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success: false,
			Message: "invalid request payload",
			Data:    err.Error(),
		})
		return err
	}

	appointments, err := cR.service.GetAppointmentByStudentID(uR)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Get appointment successfully",
		Data:    appointments,
	})
	return
}

func (cR *controllerAppointment) UpdateAppointmentStatus(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	// fmt.Println(id)
	uR := &schema.SchemaUpdateAppointmentStatus{
		ID: id,
	}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}

	appointment, err := cR.service.UpdaAppointmentStatus(uR)
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
		Message: "Update appointment status successfully",
		Data:    appointment,
	})
	return nil
}
