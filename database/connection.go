package database

import (
	"database/sql"
	"log"
	"book-restapi/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	config.LoadConfig()

	var err error
	DB, err = sql.Open("postgres", config.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	log.Println("Database connected successfully")
}