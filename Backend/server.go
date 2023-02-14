package main

import (

	"log"
	"os"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"

	routes "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/routes"
	godotenv "github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	

}

func main() {
	e := echo.New()
	host := os.Getenv("SERVER_HOST")
	port := ":" + os.Getenv("SERVER_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	client, err := ent.Open("postgres","host=" + host + " port=" + port + " user=" + db_user+" dbname=" + db_name + " password="+ db_pass)
    if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()
	routes.InitLoginRoutes(client,e)
	
	e.Logger.Fatal(e.Start(port))

}

