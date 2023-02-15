package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)


type controllerLogin struct {
	service service.ServiceLogin
}

func NewControllerLogin(service service.ServiceLogin) *controllerLogin {
	return &controllerLogin{service: service}
}

func (cR * controllerLogin) LoginUser(c echo.Context) (err error) {


	var userLogin *schema.SchemaLogin

	if err := c.Bind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success : false,
			Message : "invalid request payload",
			Data : err.Error(),
		})
		return err
	}

	userLoginInfo, err := cR.service.LoginService(userLogin)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success : false,
			Message : err.Error(),
			Data : nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success : true,
		Message : "Login successfully",
		Data : userLoginInfo,
	})
	return nil
}