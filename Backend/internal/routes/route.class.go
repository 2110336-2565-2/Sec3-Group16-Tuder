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

	// admin approve or reject class cancellation
	auditClassCancellationRoute := e.Group("/audit-class-cancellation")
	auditClassCancellationRoute.Use(mid.JWT())
	auditClassCancellationRoute.Use(mid.AdminMiddleware)
	auditClassCancellationRoute.POST("", controllerClass.AuditClassCancellation)

	// user acknowledge class rejected cancellation result
	acknowledgeClassCancellationRoute := e.Group("/acknowledge-class-cancellation")
	acknowledgeClassCancellationRoute.Use(mid.JWT())
	acknowledgeClassCancellationRoute.POST("", controllerClass.AcknowledgeClassCancellation)

	// cancel class
	cancelClassRoute := e.Group("/cancel-class")
	cancelClassRoute.Use(mid.JWT())
	cancelClassRoute.POST("", controllerClass.CancelClass)

}
