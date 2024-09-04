// name of the project

package main

import (
	"database/sql"
	"fmt"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"olympus/config"
	"olympus/internal/graphql"
)

/*
* Golang global vars
 */

var db *sql.DB

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Use the Schema variable from the graphql package
	MainSchema := graphql.MainSchema
	conf := config.New()

	pgConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Postgres.PostgresDbHost,
		conf.Postgres.PostgresDbPort,
		conf.Postgres.PostgresDbUser,
		conf.Postgres.PostgresDbPassword,
		conf.Postgres.PostgresDbName)

	//port := conf.ServerPort

	conn, err := sql.Open("postgres", pgConnStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	db = conn
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Connected to the PostgreSQL database")

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/add", addUser)
	http.HandleFunc("/users/update", updateUser)
	http.HandleFunc("/users/delete", deleteUser)

	h := handler.New(&handler.Config{
		Schema:   &MainSchema,
		Pretty:   true,
		GraphiQL: false,
	})
	http.Handle("/graphql", h)
	fmt.Printf("connect to http://localhost:8080/graphql to run graphql queries and mutations\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
