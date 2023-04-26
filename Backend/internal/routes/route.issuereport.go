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

func InitIssueReportRoutes(client *ent.Client, e *echo.Group) {
	repoIssueReport := repository.NewRepositoryIssueReport(client)
	serviceIssueReport := service.NewServiceIssueReport(repoIssueReport)
	controllerIssueReport := controller.NewControllerIssueReport(serviceIssueReport)
	middlewareInst := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))

	//e.GET("/issuereports", controllerIssueReport.GetIssueReports)
	getIssueReportsRoute := e.Group("/issuereports")
	getIssueReportsRoute.Use(middlewareInst.JWT())
	getIssueReportsRoute.Use(middlewareInst.AdminMiddleware)
	getIssueReportsRoute.GET("", controllerIssueReport.GetIssueReports)

	//e.POST("/createissuereport", controllerIssueReport.CreateIssueReport)
	createIssueReportRoute := e.Group("/createissuereport")
	createIssueReportRoute.Use(middlewareInst.JWT())
	createIssueReportRoute.POST("", controllerIssueReport.CreateIssueReport)

	//e.PUT("/issuereport/update/:id", controllerIssueReport.UpdateIssueReport)
	//e.PUT("/issuereport/updatestatus/:id", controllerIssueReport.UpdateIssueReportStatus)
	//e.DELETE("/issuereport/:id", controllerIssueReport.DeleteIssueReport)
	issueReportRoute := e.Group("/issuereport")
	issueReportRoute.Use(middlewareInst.JWT())
	issueReportRoute.Use(middlewareInst.AdminMiddleware)
	issueReportRoute.DELETE("/:id", controllerIssueReport.DeleteIssueReport)
	issueReportRoute.PUT("/updatestatus/:id", controllerIssueReport.UpdateIssueReportStatus)
	issueReportRoute.PUT("/update/:id", controllerIssueReport.UpdateIssueReport)

}
