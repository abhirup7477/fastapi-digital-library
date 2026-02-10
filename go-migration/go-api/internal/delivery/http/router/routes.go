package router

import (
	"github.com/abhirup7477/go-books-api/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, h handler.BookHandler) {
	r := engine.Group("/book-api")
	{
		r.GET("/books", h.GetAllBooks)
		r.POST("/add", h.AddBook)
		r.GET("/books/:id", h.GetBookById)
		r.PUT("/update", h.UpdateBook)
		r.DELETE("/delete/:id", h.DeleteBook)
	}
}
