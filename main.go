package main

import (
	"book-restapi/config"
	"book-restapi/database"
	"book-restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	database.RunMigrations()

	database.Connect()

	r := routes.SetupRouter()

	r.Run(":8080")
}