package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitTutorRoutes(client *ent.Client, e *echo.Group) {

	repoTutor := repository.NewRepositoryTutor(client)
	serviceTutor := service.NewServiceTutor(repoTutor)
	controllerTutor := controller.NewControllerTutor(serviceTutor)

	e.GET("/get-tutor", controllerTutor.GetTutorByUsername)
	e.GET("/get-tutors", controllerTutor.GetTutors)
	e.POST("/create-tutor", controllerTutor.CreateTutor)
	e.PUT("/update-tutor", controllerTutor.UpdateTutor)
	e.DELETE("/delete-tutor", controllerTutor.DeleteTutor)
}
