package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitCourseRoutes(client *ent.Client, e *echo.Group) {
	repoCourse := repository.NewRepositoryCourse(client)
	serviceCourse := service.NewServiceCourse(repoCourse)
	controllerCourse := controller.NewControllerCourse(serviceCourse)

	e.GET("/course", controllerCourse.GetCourseByCourseID)
	e.GET("/courses", controllerCourse.GetCourses)
	e.POST("/course", controllerCourse.CreateCourse)
	e.PUT("/course", controllerCourse.UpdateCourse)
	e.DELETE("/course", controllerCourse.DeleteCourse)
}
