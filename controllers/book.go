package controllers

import (
	"database/sql"
	"net/http"

	"book-restapi/database"
	"book-restapi/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning book data"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	query := `
		INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`
	err := database.DB.QueryRow(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, "admin").Scan(&book.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func GetBookDetail(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	query := "SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM books WHERE id = $1"
	err := database.DB.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageURL,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.CategoryID,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book details"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	query := "DELETE FROM books WHERE id = $1"
	res, err := database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
