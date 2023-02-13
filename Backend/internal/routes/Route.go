package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	"github.com/labstack/echo"
)

func SetupRouter(e *echo.Echo) {
	e.GET("/ping", controllers.PingController)
}
