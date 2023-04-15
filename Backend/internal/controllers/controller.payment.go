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
func (cR *controllerPayment) GetQRCodeForCoursePayment(c echo.Context) error {
	uR := &schema.SchemaGetQRCodeForCoursePayment{}

	if err := c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
	}

	qrCode, err := cR.service.GetQRCodeForCoursePayment(uR)

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

// GetQRCode gets QR code of a payment
func (cR *controllerPayment) GetQRCodeForTuitionFree(c echo.Context) error {
	uR := &schema.SchemaGetQRCodeForTuitionFree{}

	if err := c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
	}

	qrCode, err := cR.service.GetQRCodeForTuitionFree(uR)

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

func (cR *controllerPayment) WebhookChargeHandler(c echo.Context) error {
	uR := &schema.SchemaChargeWebhook{}

	if err := c.Bind(uR); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})
	}

	payment, err := cR.service.WebhookChargeHandler(uR)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   err,
		})
	}

	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Payment created",
		Data:    payment,
	})

	return nil
}
