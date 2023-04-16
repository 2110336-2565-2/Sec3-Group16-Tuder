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

func InitStudentRoutes(client *ent.Client, e *echo.Group) {

	repoStudent := repository.NewRepositoryStudent(client)
	serviceStudent := service.NewServiceStudent(repoStudent)
	controllerStudent := controller.NewControllerStudent(serviceStudent)
	authMiddleware := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	e.GET("/student/:username", controllerStudent.GetStudentByUsername)
	e.GET("/students", controllerStudent.GetStudents)
	e.POST("/student", controllerStudent.CreateStudent)

	//e.PUT("/student", controllerStudent.UpdateStudent)
	updateStudentRoute := e.Group("/student")
	updateStudentRoute.Use(authMiddleware.JWT())
	updateStudentRoute.Use(authMiddleware.StudentMiddleware)
	updateStudentRoute.PUT("", controllerStudent.UpdateStudent)
	
	//e.DELETE("/student", controllerStudent.DeleteStudent)
	deleteStudentRoute := e.Group("/student")
	deleteStudentRoute.Use(authMiddleware.JWT())
	deleteStudentRoute.Use(authMiddleware.AdminMiddleware)
	deleteStudentRoute.DELETE("", controllerStudent.DeleteStudent)
}
