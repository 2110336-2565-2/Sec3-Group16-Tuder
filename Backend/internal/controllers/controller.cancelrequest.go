package controllers

import (
	"net/http"

	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	services "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ControllerCancelRequest interface {
	GetCancellingRequest(echo.Context) error
	GetCancellingRequests(echo.Context) error
	CancelRequest(c echo.Context) error
	AuditRequest(c echo.Context) error
	// AcknowledgeCancelRequestCancellation(c echo.Context) error
}

type controllerCancelRequest struct {
	service services.ServiceCancelRequest
}

func NewControllerCancelRequest(s services.ServiceCancelRequest) ControllerCancelRequest {
	return &controllerCancelRequest{
		service: s,
	}
}

func (cC *controllerCancelRequest) GetCancellingRequest(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return err
	}

	CancelRequest, err := cC.service.GetCancellingRequest(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return err
	}
	c.JSON(http.StatusOK, schemas.SchemaResponses{
		Success: true,
		Message: "Get request successfully",
		Data:    CancelRequest,
	})
	return nil
}

func (cC *controllerCancelRequest) GetCancellingRequests(c echo.Context) error {
	CancelRequestes, err := cC.service.GetCancellingRequests()
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
		return err
	}
	c.JSON(http.StatusOK, schemas.SchemaResponses{
		Success: true,
		Message: "Get all requests successfully",
		Data:    CancelRequestes,
	})
	return nil
}

func (cC *controllerCancelRequest) CancelRequest(c echo.Context) error {
	request := &schemas.SchemaCancelRequest{}
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.SchemaResponses{
			Success: false,
			Message: "invalid request",
			Data:    err.Error(),
		})

		return err
	}

	err := cC.service.CancelRequest(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.SchemaResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schemas.SchemaResponses{
		Success: true,
		Message: "Cancel request has already been sent to admin",
		Data:    nil,
	})
	return nil

}

func (cC *controllerCancelRequest) AuditRequest(c echo.Context) error {

	cancelSchema := &schemas.SchemaCancelRequestApprove{}
	if err := c.Bind(cancelSchema); err != nil {
		c.JSON(http.StatusBadRequest, schemas.SchemaResponses{
			Success: false,
			Message: "invalid request",
			Data:    err.Error(),
		})
	}

	err := cC.service.AuditRequest(cancelSchema)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.SchemaResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	result := "Approved"
	if !cancelSchema.Approve {
		result = "Rejected"
	}

	c.JSON(http.StatusOK, schemas.SchemaResponses{
		Success: true,
		Message: "Cancel request has been " + result,
		Data:    nil,
	})
	return nil

}

// func (cC *controllerCancelRequest) AcknowledgeCancelRequestCancellation(c echo.Context) error {

// 	ackSchema := &schemas.SchemaUserAcknowledge{}
// 	if err := c.Bind(ackSchema); err != nil {
// 		c.JSON(http.StatusBadRequest, schemas.SchemaResponses{
// 			Success: false,
// 			Message: "invalid request",
// 			Data:    err.Error(),
// 		})
// 	}

// 	err := cC.service.AcknowledgeCancelRequestCancellation(ackSchema)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, schemas.SchemaResponses{
// 			Success: false,
// 			Message: err.Error(),
// 			Data:    nil,
// 		})
// 		return err
// 	}

// 	c.JSON(http.StatusOK, schemas.SchemaResponses{
// 		Success: true,
// 		Message: "CancelRequest cancellation has been acknowledged",
// 		Data:    nil,
// 	})
// 	return nil

// }
