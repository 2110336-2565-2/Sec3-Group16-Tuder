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

	e.GET("/appointment/:id", controllerAppointment.GetAppointmentByStudentID)
	// e.GET("/appointments", controllerAppointment.GetAppointments)
	// e.PUT("/appointment/:id", controllerAppointment.UpdateAppointment)
	// e.DELETE("/appointment/:id", controllerAppointment.DeleteAppointment)
}
