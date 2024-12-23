package database

import (
	"database/sql"
	"log"
	"os"

	"book-restapi/config"

	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB

func Connect() {
	if config.DatabaseURL == "" {
		log.Fatal("Database URL is not set")
	}

	var err error
	DB, err = sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	log.Println("Database connected successfully")
}

func RunMigrations() {
	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	migrationPath := os.Getenv("MIGRATION_PATH")
	if migrationPath == "" {
		migrationPath = "file://database/migrations"
	}

	m, err := migrate.NewWithDatabaseInstance(migrationPath, "postgres", driver)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No new migrations to apply")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		log.Println("Migrations applied successfully")
	}
}