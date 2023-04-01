package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/middlewares"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	"github.com/labstack/echo/v4"
	"os"
)

func InitReviewRoutes(client *ent.Client, e *echo.Group) {
	repoReview := repository.NewRepositoryReview(client)
	serviceReview := service.NewServiceReview(repoReview)
	controllerReview := controller.NewControllerReview(serviceReview)
	middlewareInst := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))
	e.POST("/review", controllerReview.ReviewCourse, middlewareInst.JWT())
	e.GET("/review/:courseid", controllerReview.GetRating)
}
