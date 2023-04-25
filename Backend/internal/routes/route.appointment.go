package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitAppointmentRoutes(client *ent.Client, e *echo.Group) {
	repoAppointment := repository.NewRepositoryAppointment(client)
	serviceAppointment := service.NewServiceAppointment(repoAppointment)
	controllerAppointment := controller.NewControllerAppointment(serviceAppointment)

	// add a scheduled task to the cron job

	e.GET("/matches/:id", controllerAppointment.GetMatchByID)
	e.GET("/appointments/:match_id", controllerAppointment.GetAppointmentByMatchID)
	e.PUT("/appointment/updatestatus/:id", controllerAppointment.UpdateAppointmentStatus)
}
