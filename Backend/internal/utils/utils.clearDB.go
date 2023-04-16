package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func NukeDB() {
	/*
		This function will NUKE the Database (specified in .env) and every table in it
	*/
	// retrieving vars from .env
	host := os.Getenv("SERVER_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	// create a connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable\n",
		user, password, dbname, host, dbport)
	fmt.Println(connStr)
	// roughly connect to postgres
	db, err := sql.Open("postgres",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// exec drop schema query
	_, err = db.Exec("DO $$ \nDECLARE \n  tabname RECORD; \nBEGIN \n  FOR tabname IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP \n    EXECUTE 'DROP TABLE IF EXISTS \"' || tabname.tablename || '\" CASCADE'; \n  END LOOP; \nEND $$;")
	if err != nil {
		log.Fatalf("failed executing NUKE query\n %v", err)
	}
	fmt.Println("======= public schema is NUKED =======")
}
