package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type controllerPayment struct {
	service service.ServicePayment
}

func NewControllerPayment(service service.ServicePayment) *controllerPayment {
	return &controllerPayment{service: service}
}

// GetQRCode gets QR code of a payment
func (cR *controllerPayment) GetQRCode(c echo.Context) error {
	uR := &schema.SchemaGetQRCode{}

	if err := c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
	}

	qrCode, err := cR.service.GetQRCode(uR)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
	}

	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "QR code retrieved",
		Data:    qrCode,
	})

	return nil
}
