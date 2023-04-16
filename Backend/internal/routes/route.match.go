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

func InitMatchRoutes(client *ent.Client, e *echo.Group) {
	repoMatch := repository.NewRepositoryMatch(client)
	serviceMatch := service.NewServiceMatch(repoMatch)
	controllerMatch := controller.NewControllerMatch(serviceMatch)
	middlewareInst := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))


	//e.POST("/match", controllerMatch.CreateMatch, middlewareInst.JWT())
	matchRoute := e.Group("/match")
	matchRoute.Use(middlewareInst.JWT())
	matchRoute.POST("", controllerMatch.CreateMatch)
	// e.GET("/match/:id", controllerMatch.GetMatchByMatchID)
	// e.GET("/matches", controllerMatch.GetMatches)
	// e.PUT("/match/:id", controllerMatch.UpdateMatch)
	// e.DELETE("/match/:id", controllerMatch.DeleteMatch)
}
