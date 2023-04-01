package controllers

import (
	"fmt"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type controllerReview struct {
	reviewService service.ServiceReview
}

func NewControllerReview(s service.ServiceReview) *controllerReview {
	return &controllerReview{
		reviewService: s,
	}
}

func (r controllerReview) ReviewCourse(c echo.Context) (err error) {
	cR := &schema.SchemaCreateReview{}
	if err := c.Bind(cR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return err
	}
	claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
	studentId, ok := claims["studentid"]
	if !ok {
		return fmt.Errorf("claim \"studentid\" does not exist")
	}
	cR.StudentId, err = uuid.Parse(studentId.(string))
	if err != nil {
		return fmt.Errorf("studentId from jwt token is invalid")
	}

	review, err := r.reviewService.ReviewCourseService(cR)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return err
	}
	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "review submitted successfully",
		Data:    review,
	})
	return nil
}
