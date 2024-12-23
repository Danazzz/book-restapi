package main

import (
	"book-restapi/config"
	"book-restapi/database"
	"book-restapi/routes"
)

func main() {
	config.LoadConfig()

	database.RunMigrations()

	database.Connect()

	r := routes.SetupRouter()

	r.Run(":8080")
}
