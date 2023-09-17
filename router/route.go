package router

import (
	"belajar-echo/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.POST("/books", controllers.CreateBook)
	e.GET("/books", controllers.GetBooks)
	e.GET("/books/:id", controllers.GetBookById)
	e.PATCH("/books/:id", controllers.UpdateBook)
	e.DELETE("/books/:id", controllers.DeleteBook)

}
