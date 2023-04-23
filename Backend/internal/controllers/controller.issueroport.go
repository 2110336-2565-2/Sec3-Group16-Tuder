package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type controllerIssueReport struct {
	service service.ServiceIssueReport
}

func NewControllerIssueReport(service service.ServiceIssueReport) *controllerIssueReport {
	return &controllerIssueReport{service: service}
}

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

func (cR *controllerIssueReport) UpdateIssueReport(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	uR := &schema.SchemaUpdateIssueReport{
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

	issueReport, err := cR.service.UpdateIssueReport(uR)
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
		Message: "Update issue report successfully",
		Data:    issueReport,
	})
	return nil
}

func (cR *controllerIssueReport) UpdateIssueReportStatus(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	uR := &schema.SchemaUpdateIssueReportStatus{
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

	issueReport, err := cR.service.UpdateIssueReportStatus(uR)
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
		Message: "Update issue report status successfully",
		Data:    issueReport,
	})
	return nil
}

func (cR *controllerIssueReport) DeleteIssueReport(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	uR := &schema.SchemaDeleteIssueReport{
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

	err = cR.service.DeleteIssueReport(uR)
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
		Message: "Delete issue report successfully",
	})
	return nil
}
