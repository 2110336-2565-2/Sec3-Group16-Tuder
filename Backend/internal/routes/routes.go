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
	InitStudentRoutes(c, v1)
	InitCourseRoutes(c, v1)
	InitChangePasswordRoutes(c, v1)
	InitIssueReportRoutes(c, v1)
	InitCancelRequestRoutes(c, v1)
	InitReviewRoutes(c, v1)
	InitMatchRoutes(c, v1)
	InitAppointmentRoutes(c, v1)

}
