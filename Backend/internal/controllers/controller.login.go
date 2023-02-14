package controllers

import(
	"net/http"
	echo "github.com/labstack/echo/v4"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
)


type controllerLogin struct {
	service service.ServiceLogin
}

func NewControllerLogin(service service.ServiceLogin) *controllerLogin {
	return &controllerLogin{service: service}
}

func (c * controllerLogin) LoginUser(ctx echo.Context) (err error) {

	var userLogin *schema.SchemaLogin

	if err := ctx.Bind(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}

	userLogin, err = c.service.LoginService(userLogin)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}

	ctx.JSON(http.StatusOK, userLogin)
	return nil
}