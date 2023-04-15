package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitPaymentRoutes(client *ent.Client, e *echo.Group) {
	repoPayment := repository.NewRepositoryPayment(client)
	servicePayment := service.NewServicePayment(repoPayment)
	controllerPayment := controller.NewControllerPayment(servicePayment)

	e.POST("/payment/getQRCodeForCoursePayment", controllerPayment.GetQRCodeForCoursePayment)
	e.POST("/payment/getQRCodeForTuitionFree", controllerPayment.GetQRCodeForTuitionFree)
	e.POST("/payment/webhookChargeHandler", controllerPayment.WebhookChargeHandler)
}
