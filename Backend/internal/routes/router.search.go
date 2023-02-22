package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitSearchRoutes(c *ent.Client, e *echo.Group) {

	repoSearch := repository.NewRepositoryLogin(client)
	serviceSearch := service.NewServiceLogin(repoSearch)
	controllerSearch := controller.NewControllerLogin(serviceSearch)

	e.POST("/search", controllerSearch.SearchContent)

}
