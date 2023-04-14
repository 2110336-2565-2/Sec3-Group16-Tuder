package routes

import (
	"os"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/middlewares"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitTutorRoutes(client *ent.Client, e *echo.Group) {

	repoTutor := repository.NewRepositoryTutor(client)
	serviceTutor := service.NewServiceTutor(repoTutor)
	controllerTutor := controller.NewControllerTutor(serviceTutor)
	middlewareInst := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	e.GET("/tutor/:username", controllerTutor.GetTutorByUsername)
	e.GET("/tutorID/:id", controllerTutor.GetTutorByID)
	e.GET("/tutors", controllerTutor.GetTutors)
	e.POST("/tutor", controllerTutor.CreateTutor)
	e.PUT("/tutor", controllerTutor.UpdateTutor)
	e.DELETE("/tutor", controllerTutor.DeleteTutor)
	e.PUT("/tutor/schedule", controllerTutor.UpdateSchedule)
	e.GET("/tutor/schedule", controllerTutor.GetTutorSchedule)
	e.GET("/tutor/schedule/:courseId", controllerTutor.GetTutorScheduleByCourseId)
	e.GET("/tutor/:username/reviews", controllerTutor.GetTutorReviews)
	e.GET("/tutor/courses", controllerTutor.GetTutorCourses, middlewareInst.JWT())
}
