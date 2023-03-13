package controllers

import (
	"net/http"
	"github.com/golang-jwt/jwt/v4"


	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type controllerStudent struct {
	studentService service.ServiceStudent
}

func NewControllerStudent(s service.ServiceStudent) *controllerStudent {
	return &controllerStudent{studentService: s}
}

func (cS *controllerStudent) GetStudentByUsername(c echo.Context) (err error) {
	uR := &schema.SchemaGetStudent{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "Invalid request",
			Error:   err,
		})
		return
	}

	student, err := cS.studentService.GetStudentByUsername(uR)
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
		Message: "Get student successfully",
		Data:    student,
	})
	return nil
}

func (cS *controllerStudent) GetStudents(c echo.Context) (err error) {
	students, err := cS.studentService.GetStudents()
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
	}
	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Get students successfully",
		Data:    students,
	})
	return nil
}


func (cS *controllerStudent) CreateStudent(c echo.Context) (err error) {
	uR := &schema.SchemaCreateStudent{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "Invalid request",
			Error:   err,
		})
		return
	}

	student, err := cS.studentService.CreateStudent(uR)
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
		Message: "Create student successfully",
		Data:    student,
	})
	return nil
}

func (cS *controllerStudent) UpdateStudent(c echo.Context) (err error) {
	uR := &schema.SchemaUpdateStudent{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "Invalid request",
			Error:   err,
		})
		return
	}
	uR.Username = c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["username"].(string)
	student, err := cS.studentService.UpdateStudent(uR)
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
		Message: "Update student successfully",
		Data:    student,
	})
	return nil
}

func (cS *controllerStudent) DeleteStudent(c echo.Context) (err error) {
	uR := &schema.SchemaDeleteStudent{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "Invalid request",
			Error:   err,
		})
		return
	}

	err = cS.studentService.DeleteStudent(uR)
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
		Message: "Delete student successfully",
	})
	return nil
}
