package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/middlewares"
	"os"
)

func InitAdminTuitionFeeRoutes(client *ent.Client, e *echo.Group) {
	repoAdminTuitionFee := repository.NewRepositoryAdminTuitionFee(client)
	serviceAdminTuitionFee := service.NewServiceAdminTuitionFee(repoAdminTuitionFee)
	controllerAdminTuitionFee := controller.NewControllerAdminTuitionFee(serviceAdminTuitionFee)
	middlewareInst := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	//e.GET("/admintuitionfees", controllerAdminTuitionFee.GetAdminTuitionFees)
	adminTuitionFeeRoute := e.Group("/admintuitionfees")
	adminTuitionFeeRoute.Use(middlewareInst.JWT())
	adminTuitionFeeRoute.Use(middlewareInst.AdminMiddleware)
	adminTuitionFeeRoute.GET("", controllerAdminTuitionFee.GetAdminTuitionFees)
}
