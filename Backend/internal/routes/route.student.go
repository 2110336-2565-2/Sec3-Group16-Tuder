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

func InitStudentRoutes(client *ent.Client, e *echo.Group) {

	repoStudent := repository.NewRepositoryStudent(client)
	serviceStudent := service.NewServiceStudent(repoStudent)
	controllerStudent := controller.NewControllerStudent(serviceStudent)
	authMiddleware := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	e.GET("/student", controllerStudent.GetStudentByUsername)
	e.GET("/students", controllerStudent.GetStudents)
	e.POST("/student", controllerStudent.CreateStudent)
	e.PUT("/student", controllerStudent.UpdateStudent, authMiddleware.JWT())
	e.DELETE("/student", controllerStudent.DeleteStudent, authMiddleware.JWT())
}
