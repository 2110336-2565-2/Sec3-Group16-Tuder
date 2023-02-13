package main

import (
	"log"
	"os"
	routes "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	e := echo.New()
	routes.InitUserRoutes(e)
	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))

}
