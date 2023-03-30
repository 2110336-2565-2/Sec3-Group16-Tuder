package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitIssueReportRoutes(client *ent.Client, e *echo.Group) {
	repoIssueReport := repository.NewRepositoryIssueReport(client)
	serviceIssueReport := service.NewServiceIssueReport(repoIssueReport)
	controllerIssueReport := controller.NewControllerIssueReport(serviceIssueReport)

	// e.GET("/course/:id", controllerCourse.GetCourseByCourseID)
	e.GET("/issuereports", controllerIssueReport.GetIssueReports)
	e.POST("/createissuereport", controllerIssueReport.CreateIssueReport)
	// e.PUT("/course/:id", controllerCourse.UpdateCourse)
	// e.DELETE("/course/:id", controllerCourse.DeleteCourse)
	// e.POST("/coursesearch", controllerCourse.SearchContent)
}
