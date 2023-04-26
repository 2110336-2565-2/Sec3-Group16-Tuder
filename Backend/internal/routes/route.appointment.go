package routes

import (
	"os"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	middlewares "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/middlewares"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitAppointmentRoutes(client *ent.Client, e *echo.Group) {
	repoAppointment := repository.NewRepositoryAppointment(client)
	serviceAppointment := service.NewServiceAppointment(repoAppointment)
	controllerAppointment := controller.NewControllerAppointment(serviceAppointment)
	middlewareInst := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	// add a scheduled task to the cron job

	e.GET("/matches/:id", controllerAppointment.GetMatchByID)
	e.GET("/appointments/:match_id", controllerAppointment.GetAppointmentByMatchID)
	updateStatusRoute := e.Group("/appointment/updatestatus/:id")
	updateStatusRoute.Use(middlewareInst.JWT())
	updateStatusRoute.Use(middlewareInst.StudentMiddleware)
	updateStatusRoute.PUT("", controllerAppointment.UpdateAppointmentStatus)
}
