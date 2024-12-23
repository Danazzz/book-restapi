package main

import (
	"book-restapi/database"
	"book-restapi/routes"
)

func main() {
	database.Connect()

	r := routes.SetupRouter()

	r.Run(":8080")
}