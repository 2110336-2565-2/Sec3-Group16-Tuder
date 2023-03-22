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

func InitClassRoutes(client *ent.Client, e *echo.Group) {
	repoClass := repository.NewRepositoryClass(client)
	serviceClass := service.NewServiceClass(repoClass)
	controllerClass := controller.NewControllerClass(serviceClass)
	mid := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	// get cancelling classes
	getCancellingClassesRoute := e.Group("/cancelling-classes")
	getCancellingClassesRoute.Use(mid.JWT())
	getCancellingClassesRoute.Use(mid.AdminMiddleware)
	getCancellingClassesRoute.GET("", controllerClass.GetCancellingClasses)

	// approve class cancellation
	approveClassCancellationRoute := e.Group("/approve-class-cancellation")
	approveClassCancellationRoute.Use(mid.JWT())
	approveClassCancellationRoute.Use(mid.AdminMiddleware)
	approveClassCancellationRoute.POST("", controllerClass.ApproveClassCancellation)

	// cancel class
	cancelClassRoute := e.Group("/cancel-class")
	cancelClassRoute.Use(mid.JWT())
	cancelClassRoute.POST("", controllerClass.CancelClass)

}
