package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

func InitRegisterRoute(c *ent.Client, e *echo.Group) {
	repoRegister := repository.NewRepositoryRegister(c)
	serviceRegister := services.NewServiceRegister(repoRegister)
	controllerRegister := controller.NewControllerRegister(serviceRegister)

	e.POST("/register", controllerRegister.RegisterUser)
}