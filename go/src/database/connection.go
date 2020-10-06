package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	
)

const (
	// Initialize connection constants.
	HOST     = "localhost"
	DATABASE = "estudiantes_db"
	USER     = "postgres"
	PASSWORD = "password"
)



func GetConnection() *sql.DB {
	/* var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)	 */
	connStr := "postgres://postgres:password@localhost/estudiantes_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		/*  */
	}

 	/* if err != nil {		
		panic(err)	
	} */
	
	return db
}
