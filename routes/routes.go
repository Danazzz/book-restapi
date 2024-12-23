package routes

import (
	"book-restapi/controllers"
	"book-restapi/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API untuk login dan register
	r.POST("/api/users/login", controllers.Login)
	r.POST("/api/users/register", controllers.Register)

	// API dengan middleware JWT
	api := r.Group("/api")
	api.Use(middlewares.JWTMiddleware()) // Middleware untuk autentikasi JWT
	{
		api.GET("/books", controllers.GetBooks)
		api.POST("/books", controllers.AddBook)
		api.GET("/books/:id", controllers.GetBookDetail)
		api.DELETE("/books/:id", controllers.DeleteBook)

		api.GET("/categories", controllers.GetCategories)
		api.POST("/categories", controllers.AddCategory)
		api.GET("/categories/:id", controllers.GetCategoryDetail)
		api.DELETE("/categories/:id", controllers.DeleteCategory)
	}

	return r
}