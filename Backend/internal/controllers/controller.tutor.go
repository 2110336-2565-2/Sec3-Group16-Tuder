package controllers

import (
	// "github.com/golang-jwt/jwt/v4"

	"net/http"

	"github.com/google/uuid"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type ControllerTutor interface {
	CreateTutor(c echo.Context) error
	GetTutors(c echo.Context) error
	GetTutorByUsername(c echo.Context) error
	GetTutorByID(c echo.Context) error
	GetTutorSchedule(c echo.Context) error
	UpdateTutor(c echo.Context) error
	UpdateSchedule(c echo.Context) error
	DeleteTutor(c echo.Context) error
	GetTutorReviews(c echo.Context) error
}

type controllerTutor struct {
	service service.ServiceTutor
}

func NewControllerTutor(service service.ServiceTutor) *controllerTutor {
	return &controllerTutor{service: service}
}

func (cR *controllerTutor) GetTutorByUsername(c echo.Context) (err error) {
	username := c.Param("username")
	uR := &schema.SchemaGetTutor{
		Username: username,
	}
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

func (cR *controllerTutor) GetTutorScheduleByCourseId(c echo.Context) error {
	courseId, _ := uuid.Parse(c.Param("courseId"))
	uR := &schema.SchemaGetTutorScheduleByCourseId{
		Course_id: courseId,
	}

	tutorSchedule, err := cR.service.GetTutorScheduleByCourseId(uR)

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
		Message: "Get tutor schedule successfully",
		Data:    tutorSchedule,
	})
	return nil

}
func (cR *controllerTutor) GetTutorByID(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	uR := &schema.SchemaGetTutorByID{
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

	tutor, err := cR.service.GetTutorByID(uR)

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
	// uR.Username = c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["username"].(string)
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

func (cR *controllerTutor) UpdateSchedule(c echo.Context) (err error) {
	s := &schema.SchemaUpdateSchedule{}
	if err = c.Bind(s); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}
	// s.Username = c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["username"].(string)
	schedule, err := cR.service.UpdateTutorSchedule(s)

	// success
	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Schedule edited successfully",
		Data:    schedule,
	})
	return nil
}

func (cR controllerTutor) GetTutorSchedule(c echo.Context) (err error) {
	s := &schema.SchemaGetSchedule{}
	if err = c.Bind(s); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}
	// s.Username = c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["username"].(string)
	schedule, err := cR.service.GetTutorSchedule(s)
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
		Message: "Get Schedule successfully",
		Data:    schedule,
	})
	return
}

func (cR controllerTutor) GetTutorReviews(c echo.Context) (err error) {
	username := c.Param("username")
	r := &schema.SchemaGetReviews{
		Username: username,
	}
	if err = c.Bind(r); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}
	// r.Username = c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["username"].(string)
	reviews, err := cR.service.GetTutorReviews(r)
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
		Message: "Get Reviews successfully",
		Data:    reviews,
	})
	return
}
