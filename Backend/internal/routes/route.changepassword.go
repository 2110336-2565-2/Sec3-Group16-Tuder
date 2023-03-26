package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitChangePasswordRoutes(client *ent.Client, e *echo.Group) {
	repoChangePassword := repository.NewRepositoryChangePassword(client)
	serviceChangePassword := service.NewServiceChangePassword(repoChangePassword)
	controllerChangePassword := controller.NewControllerChangePassword(serviceChangePassword)

	e.POST("/changepassword", controllerChangePassword.ChangePassword)
	e.POST("/changepassword-check", controllerChangePassword.CheckPassword)

}
