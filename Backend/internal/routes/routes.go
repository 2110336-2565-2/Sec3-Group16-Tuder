package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/labstack/echo/v4"
)

func InitRoutes(c *ent.Client, e *echo.Echo) {
	v1 := e.Group("/api/v1")
	InitLoginRoutes(c, v1)
	InitRegisterRoute(c, v1)
	InitTutorRoutes(c, v1)
}
