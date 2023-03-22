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

func InitCancelClassRoutes(client *ent.Client, e *echo.Group) {
	repoCancelClass := repository.NewRepositoryCancelClass(client)
	serviceCancelClass := service.NewServiceCancelClass(repoCancelClass)
	controllerCancelClass := controller.NewControllerCancelClass(serviceCancelClass)
	mid := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	cancelClassRoute := e.Group("/cancelclass")
	cancelClassRoute.Use(mid.JWT())

	cancelClassRoute.POST("", controllerCancelClass.CancelClass)

}
