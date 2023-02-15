package controllers

import (
	"fmt"
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type controllerRegister struct {
	studentService service.ServiceStudentRegister
	tutorService   service.ServiceTutorRegister
}

func NewControllerRegister(sS service.ServiceStudentRegister, sT service.ServiceTutorRegister) *controllerRegister {
	return &controllerRegister{tutorService: sT, studentService: sS}
}

func (cR *controllerRegister) RegisterUser(c echo.Context) (err error) {

	type Role struct {
		As string `json:"as,omitempty"`
	}
	uRole := &Role{}
	if err = c.Bind(uRole); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaRegisterResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}

	uR := &schema.SchemaRegister{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaRegisterResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}

	if uRole.As == "student" {
		if err = cR.studentService.RegisterStudentService(uR); err != nil {
			c.JSON(http.StatusBadRequest, schema.SchemaRegisterResponse{
				Success: false,
				Message: fmt.Sprint(err),
				Error:   err,
			})
			return
		}
	}
	if uRole.As == "tutor" {
		if err = cR.tutorService.RegisterTutorService(uR); err != nil {
			c.JSON(http.StatusBadRequest, schema.SchemaRegisterResponse{
				Success: false,
				Message: fmt.Sprint(err),
				Error:   err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, schema.SchemaRegisterResponse{
		Success: true,
		Message: "Register successfully",
	})
	return nil
}
