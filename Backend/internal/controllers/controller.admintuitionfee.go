package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"

	"github.com/labstack/echo/v4"
)

type controllerAdminTuitionFee struct {
	service service.ServiceAdminTuitionFee
}

func NewControllerAdminTuitionFee(service service.ServiceAdminTuitionFee) *controllerAdminTuitionFee {
	return &controllerAdminTuitionFee{service: service}
}

func (cR *controllerAdminTuitionFee) GetAdminTuitionFees(c echo.Context) (err error) {
	adminTuitionFeeResult, err := cR.service.GetAdminTuitionFees()

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schema.SchemaAdminTuitionFeeResponse{
		Success: true,
		Message: "Search successfully",
		Data:    adminTuitionFeeResult,
	})
	return
}
