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

func InitCourseRoutes(client *ent.Client, e *echo.Group) {
	repoCourse := repository.NewRepositoryCourse(client)
	serviceCourse := service.NewServiceCourse(repoCourse)
	controllerCourse := controller.NewControllerCourse(serviceCourse)
	middlewareInst := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	e.GET("/course/:id", controllerCourse.GetCourseByCourseID)
	e.GET("/courses", controllerCourse.GetCourses)
	e.POST("/coursesearch", controllerCourse.SearchContent)

	//e.PUT("/course/status/:id", controllerCourse.UpdateCourseStatus)
	//e.DELETE("/course/:id", controllerCourse.DeleteCourse)
	//e.POST("/course", controllerCourse.CreateCourse)
	//e.PUT("/course/:id", controllerCourse.UpdateCourse)
	courseRoute := e.Group("/course")
	courseRoute.Use(middlewareInst.JWT())
	// courseRoute.Use(middlewareInst.TutorMiddleware)
	courseRoute.PUT("/status/:id", controllerCourse.UpdateCourseStatus)
	courseRoute.DELETE("/:id", controllerCourse.DeleteCourse)
	courseRoute.POST("", controllerCourse.CreateCourse)
	courseRoute.PUT("/:id", controllerCourse.UpdateCourse)
}
