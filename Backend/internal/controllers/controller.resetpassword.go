package controllers

import (
	"fmt"
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

type ControllerResetPassword interface {
	ForgetPassword(c echo.Context) error
	ResetPassword(c echo.Context) error
}
type controllerResetPassword struct {
	service service.ServiceResetPassword
}

func NewControllerResetPassword(s service.ServiceResetPassword) *controllerResetPassword {
	return &controllerResetPassword{service: s}
}

func (cR *controllerResetPassword) ForgetPassword(c echo.Context) error {
	// Parse request
	req := &schema.SchemaResetPassword{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})

		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	// Reset password
	if err := cR.service.SentVerificationEmailService(req); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: fmt.Sprint(err),
			Error:   err,
		})
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to reset password")
	}
	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Resetpassword successfully",
	})
	return c.NoContent(http.StatusOK)
}

func (cR *controllerResetPassword) ResetPassword(c echo.Context) error {
	// Parse request
	req := &schema.SchemaNewPassword{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: "invalid request payload",
			Error:   err,
		})

		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	// Reset password
	if err := cR.service.ResetPasswordService(req); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaErrorResponse{
			Success: false,
			Message: fmt.Sprint(err),
			Error:   err,
		})
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to reset password")
	}
	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success: true,
		Message: "Resetpassword successfully",
	})
	return c.NoContent(http.StatusOK)
}
