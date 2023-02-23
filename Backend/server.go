package main

import (
	"context"
	"log"
	"os"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/migrate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/datas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/middlewares"
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
	db_port := os.Getenv("DB_PORT")

	client, err := ent.Open("postgres", "host="+host+" port="+db_port+" user="+db_user+" dbname="+db_name+" password="+db_pass+" sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()


	if err := client.Schema.Create(context.Background(),

		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// test must reset db as always
	datas.InsertData(client)

	routes.InitRoutes(client, e)
	e.Use(middlewares.CorsMiddleware)
	e.Logger.Fatal(e.Start(port))

}
