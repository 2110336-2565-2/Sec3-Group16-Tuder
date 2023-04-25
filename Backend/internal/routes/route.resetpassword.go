package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/middlewares"
	"os"
)

func InitResetPasswordRoute(c *ent.Client, e *echo.Group) {
	repoResetPassword := repository.NewRepositoryResetPassword(c)
	serviceResetPassword := services.NewServiceResetPassword(repoResetPassword)
	controllerResetPassword := controller.NewControllerResetPassword(serviceResetPassword)
	middlewareInst := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	e.POST("/forget-password", controllerResetPassword.ForgetPassword)

	//e.POST("/reset-password", controllerResetPassword.ResetPassword)
	resetPasswordRoute := e.Group("/reset-password")
	resetPasswordRoute.Use(middlewareInst.JWT())
	resetPasswordRoute.POST("", controllerResetPassword.ResetPassword)
}
