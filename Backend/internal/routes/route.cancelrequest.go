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

	// get cancelling Request by id
	getCancellingRequestRoute := e.Group("/cancelling-request/:id")
	getCancellingRequestRoute.Use(mid.JWT())
	getCancellingRequestRoute.Use(mid.AdminMiddleware)
	getCancellingRequestRoute.GET("", controllerCancelRequest.GetCancellingRequest)

	// get cancelling requests
	getCancellingRequestsRoute := e.Group("/cancelling-requests")
	getCancellingRequestsRoute.Use(mid.JWT())
	getCancellingRequestsRoute.Use(mid.AdminMiddleware)
	getCancellingRequestsRoute.GET("", controllerCancelRequest.GetCancellingRequests)

	// admin approve or reject request
	auditMatchCancellationRoute := e.Group("/audit-request")
	auditMatchCancellationRoute.Use(mid.JWT())
	auditMatchCancellationRoute.Use(mid.AdminMiddleware)
	auditMatchCancellationRoute.POST("", controllerCancelRequest.AuditRequest)

	// // user acknowledge request rejected cancellation result
	// acknowledgeCancelRequestCancellationRoute := e.Group("/acknowledge-CancelRequest-cancellation")
	// acknowledgeCancelRequestCancellationRoute.Use(mid.JWT())
	// acknowledgeCancelRequestCancellationRoute.POST("", controllerCancelRequest.AcknowledgeCancelRequestCancellation)

	// cancel request
	cancelRequestRoute := e.Group("/cancel-request")
	cancelRequestRoute.Use(mid.JWT())
	cancelRequestRoute.POST("", controllerCancelRequest.CancelRequest)

}
