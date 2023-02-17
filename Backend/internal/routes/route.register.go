package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
)

func InitRegisterRoute(c *ent.Client, e *echo.Group) {
	repoStudentRegister := repository.NewRepositoryStudentRegister(c)
	repoTutorRegister := repository.NewRepositoryTutorRegister(c)
	serviceStudentRegister := services.NewServiceStudentRegister(repoStudentRegister)
	serviceTutorRegister := services.NewServiceTutorRegister(repoTutorRegister)
	controllerRegister := controller.NewControllerRegister(serviceStudentRegister, serviceTutorRegister)

	e.POST("/register", controllerRegister.RegisterUser)
}
