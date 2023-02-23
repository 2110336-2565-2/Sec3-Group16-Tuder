package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitCourseSearchRoutes(c *ent.Client, e *echo.Group) {

	repoCourseSearch := repository.NewRepositoryCourseSearch(c)
	serviceCourseSearch := service.NewServiceCourseSearch(repoCourseSearch)
	controllerCourseSearch := controller.NewControllerCourseSearch(serviceCourseSearch)

	e.POST("/coursesearch", controllerCourseSearch.SearchContent)
	e.POST("/courses", controllerCourseSearch.GetAllCourse)
	
}
