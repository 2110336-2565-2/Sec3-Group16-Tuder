package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/migrate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/middlewares"
	routes "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/routes"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
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
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// test must reset db as always
	ctx := context.Background()

	client.Tutor.Delete().Exec(ctx)
	client.Student.Delete().Exec(ctx)
	client.User.Delete().Exec(ctx)



	ps, _ := utils.HashPassword("brightHeemen")
	user1, err := client.User.Create().
		SetUsername("hee").
		SetPassword(ps).
		SetAddress("a").
		SetEmail("a").
		SetPhone("0").
		SetFirstName("Bright").
		SetLastName("Jukjeejid").
		SetGender("male").
		SetBirthDate(time.Now()).
		SetRole(user.RoleStudent).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	user2, err := client.User.Create().
		SetUsername("bighee").
		SetPassword(ps).
		SetAddress("b").
		SetEmail("b").
		SetPhone("1").
		SetFirstName("Bright").
		SetLastName("Jukjeejid").
		SetGender("female").
		SetBirthDate(time.Now()).
		SetRole(user.RoleTutor).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	fmt.Println(user1)
	fmt.Println(user2)

	client.Student.Create().
		SetUserID(user1.ID).
		Save(context.Background())

	client.Tutor.Create().
		SetUserID(user2.ID).
		SetCitizenID("1234567890123").
		Save(context.Background())

	routes.InitRoutes(client, e)
	e.Use(middlewares.CorsMiddleware)
	e.Logger.Fatal(e.Start(port))

}
