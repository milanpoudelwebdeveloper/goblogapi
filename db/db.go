package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

var err error

func InitiDB() {
	connStr := "user=postgres dbname=myblogserver password=milanpoudel host=localhost sslmode=disable"
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid.", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database.", err)
	}
	log.Println("Connected to the database.")
}
