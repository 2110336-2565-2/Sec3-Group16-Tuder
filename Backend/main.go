package main

import (
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"os"
)

func main() {
	godotenv.Load()
	port := os.Getenv("APP_PORT")
	url := os.Getenv("APP_URL")
	e := echo.New()
	routes.SetupRouter(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", url, port)))
}
