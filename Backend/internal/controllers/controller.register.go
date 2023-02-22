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

	
	uR := &schema.SchemaRegister{}
	if err := c.Bind(&uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaRegisterResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return err
	}


	role  := uR.Role

	token := ""
	if role == "student" {
		userLoginInfo, err := cR.studentService.RegisterStudentService(uR)
		if err != nil {

			c.JSON(http.StatusBadRequest, schema.SchemaRegisterResponse{
				Success: false,
				Message: fmt.Sprint(err),
				Error:   err,
			})
			return err
		}
		token = userLoginInfo.Token
	}
	if role == "tutor" {
		userLoginInfo, err := cR.tutorService.RegisterTutorService(uR)
		if err != nil {

			c.JSON(http.StatusBadRequest, schema.SchemaRegisterResponse{
				Success: false,
				Message: fmt.Sprint(err),
				Error:   err,
			})
			return err
		}
		token = userLoginInfo.Token
	}

	c.JSON(http.StatusOK, schema.SchemaRegisterResponse{
		Success: true,
		Message: "Register successfully",
		Token:   token,
	})
	return nil
}
