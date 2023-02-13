package main

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routes.SetupRouter(e)
	e.Logger.Fatal(e.Start(":10000"))
}
