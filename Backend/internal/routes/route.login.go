package routes

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	controller "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/controllers"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
)


func InitLoginRoutes(client *ent.Client,e *echo.Echo){
	

	repoLogin := repository.NewRepositoryLogin(client)
	serviceLogin := service.NewServiceLogin(repoLogin)
	controllerLogin := controller.NewControllerLogin(serviceLogin)


	login := e.Group("/api/v1")
	login.POST("/login", controllerLogin.LoginUser)

}
