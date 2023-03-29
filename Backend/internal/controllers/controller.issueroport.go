package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"

	// "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type controllerIssueReport struct {
	service service.ServiceIssueReport
}

func NewControllerIssueReport(service service.ServiceIssueReport) *controllerIssueReport {
	return &controllerIssueReport{service: service}
}

// func (cR *controllerCourse) GetCourseByCourseID(c echo.Context) (err error) {
// 	id, _ := uuid.Parse(c.Param("id"))
// 	uR := &schema.SchemaGetCourse{
// 		ID: id,
// 	}

// 	course, err := cR.service.GetCourseByID(uR)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
// 			Success: false,
// 			Message: err.Error(),
// 			Error:   err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, schema.SchemaResponses{
// 		Success: true,
// 		Message: "Get course successfully",
// 		Data:    course,
// 	})
// 	return nil
// }

func (cR *controllerIssueReport) GetIssueReports(c echo.Context) (err error) {
	issueReportResult, err := cR.service.GetIssueReports()

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schema.SchemaReportResponse{
		Success: true,
		Message: "Search successfully",
		Data:    issueReportResult,
	})
	return
}

func (cR *controllerIssueReport) CreateIssueReport(c echo.Context) (err error) {
	uR := &schema.SchemaCreateIssueReport{}
	if err = c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
		return
	}

	issueReport, err := cR.service.CreateIssueReport(uR)
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
		Message: "Create issue report successfully",
		Data:    issueReport,
	})
	return nil
}

// func (cR *controllerCourse) UpdateCourse(c echo.Context) (err error) {
// 	id, _ := uuid.Parse(c.Param("id"))
// 	uR := &schema.SchemaUpdateCourse{
// 		ID: id,
// 	}
// 	if err = c.Bind(uR); err != nil {
// 		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
// 			Success: false,
// 			Message: "invalid request payload",
// 			Error:   err,
// 		})
// 		return
// 	}

// 	course, err := cR.service.UpdateCourse(uR)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
// 			Success: false,
// 			Message: err.Error(),
// 			Error:   err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, schema.SchemaResponses{
// 		Success: true,
// 		Message: "Update course successfully",
// 		Data:    course,
// 	})
// 	return nil
// }

// func (cR *controllerCourse) DeleteCourse(c echo.Context) (err error) {
// 	id, _ := uuid.Parse(c.Param("id"))
// 	uR := &schema.SchemaDeleteCourse{
// 		ID: id,
// 	}
// 	if err = c.Bind(uR); err != nil {
// 		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
// 			Success: false,
// 			Message: "invalid request payload",
// 			Error:   err,
// 		})
// 		return
// 	}

// 	err = cR.service.DeleteCourse(uR)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
// 			Success: false,
// 			Message: err.Error(),
// 			Error:   err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, schema.SchemaResponses{
// 		Success: true,
// 		Message: "Delete course successfully",
// 	})
// 	return nil
// }
