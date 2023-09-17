package controllers

import (
	"belajar-echo/config"
	"belajar-echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	var bookReq models.BookReq

	if err := c.Bind(&bookReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var book models.Book

	book.Title = bookReq.Title
	book.Author = bookReq.Author
	book.PublishYear = bookReq.PublishYear
	book.ISBN = bookReq.ISBN
	book.Genre = bookReq.Genre

	if err := config.DB.Create(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success creation book",
		"book":    book,
	})
}

func GetBooks(c echo.Context) error {
	var books []models.Book

	config.DB.Find(&books)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book found successfully",
		"books":   books,
	})
}

func GetBookById(c echo.Context) error {
	id := c.Param("id")

	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book found successfully",
		"book":    book,
	})
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")

	// Cari buku berdasarkan ID
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Book not found",
		})
	}

	// Bind request payload ke struct BookReq
	var bookReq models.BookReq
	if err := c.Bind(&bookReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Payload invalid",
		})
	}

	// Update informasi buku
	if bookReq.Title != "" {
		book.Title = bookReq.Title
	}
	if bookReq.Author != "" {
		book.Author = bookReq.Author
	}
	if bookReq.PublishYear != 0 {
		book.PublishYear = bookReq.PublishYear
	}
	if bookReq.ISBN != "" {
		book.ISBN = bookReq.ISBN
	}
	if bookReq.Genre != "" {
		book.Genre = bookReq.Genre
	}

	// Simpan perubahan ke database
	if err := config.DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book updated successfully",
		"book":    book,
	})
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Book not found",
		})
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete book",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book deleted successfully",
	})
}
