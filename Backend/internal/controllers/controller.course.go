package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type controllerCourse struct {
	service service.ServiceCourse
}

func NewControllerCourse(service service.ServiceCourse) *controllerCourse {
	return &controllerCourse{service: service}
}

func (cR *controllerCourse) GetCourseByCourseID(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	uR := &schema.SchemaGetCourse{
		ID: id,
	}

	course, err := cR.service.GetCourseByID(uR)
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
		Message: "Get course successfully",
		Data:    course,
	})
	return nil
}

func (cR *controllerCourse) GetCourses(c echo.Context) (err error) {
	courses, err := cR.service.GetCourses()
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
		Message: "Get courses successfully",
		Data:    courses,
	})
	return nil
}

func (cR *controllerCourse) CreateCourse(c echo.Context) (err error) {
	uR := &schema.SchemaCreateCourse{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}

	course, err := cR.service.CreateCourse(uR)
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
		Message: "Create course successfully",
		Data:    course,
	})
	return nil
}

func (cR *controllerCourse) UpdateCourse(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	uR := &schema.SchemaUpdateCourse{
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

	course, err := cR.service.UpdateCourse(uR)
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
		Message: "Update course successfully",
		Data:    course,
	})
	return nil
}

func (cR *controllerCourse) DeleteCourse(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	uR := &schema.SchemaDeleteCourse{
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

	err = cR.service.DeleteCourse(uR)
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
		Message: "Delete course successfully",
	})
	return nil
}
