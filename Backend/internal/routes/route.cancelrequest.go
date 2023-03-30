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

func InitCancelRequestRoutes(client *ent.Client, e *echo.Group) {
	repoCancelRequest := repository.NewRepositoryCancelRequest(client)
	serviceCancelRequest := service.NewServiceCancelRequest(repoCancelRequest)
	controllerCancelRequest := controller.NewControllerCancelRequest(serviceCancelRequest)
	mid := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	// get cancelling CancelRequestes
	getCancellingCancelRequestesRoute := e.Group("/cancelling-requests")
	getCancellingCancelRequestesRoute.Use(mid.JWT())
	getCancellingCancelRequestesRoute.Use(mid.AdminMiddleware)
	getCancellingCancelRequestesRoute.GET("", controllerCancelRequest.GetCancellingRequests)

	// admin approve or reject CancelRequest cancellation
	auditCancelRequestCancellationRoute := e.Group("/audit-request")
	auditCancelRequestCancellationRoute.Use(mid.JWT())
	auditCancelRequestCancellationRoute.Use(mid.AdminMiddleware)
	auditCancelRequestCancellationRoute.POST("", controllerCancelRequest.AuditRequest)

	// // user acknowledge CancelRequest rejected cancellation result
	// acknowledgeCancelRequestCancellationRoute := e.Group("/acknowledge-CancelRequest-cancellation")
	// acknowledgeCancelRequestCancellationRoute.Use(mid.JWT())
	// acknowledgeCancelRequestCancellationRoute.POST("", controllerCancelRequest.AcknowledgeCancelRequestCancellation)

	// cancel CancelRequest
	cancelCancelRequestRoute := e.Group("/cancel-request")
	cancelCancelRequestRoute.Use(mid.JWT())
	cancelCancelRequestRoute.POST("", controllerCancelRequest.CancelRequest)

}
