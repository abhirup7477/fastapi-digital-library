package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/abhirup7477/go-books-api/internal/delivery/http/Response"
	"github.com/abhirup7477/go-books-api/internal/delivery/http/dto"
	"github.com/abhirup7477/go-books-api/internal/domain"
	"github.com/abhirup7477/go-books-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	uc usecase.BookUsecase
}

func NewHandler(uc *usecase.BookUsecase) *BookHandler {
	return &BookHandler{uc: *uc}
}

func (h *BookHandler) UsecaseRequirements(c *gin.Context) (context.Context, string) {
	ctx := c.Request.Context()
	email := c.Request.Header.Get(("email"))
	return ctx, email
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var bookModel dto.BookModel

	if err := c.ShouldBindJSON(&bookModel); err != nil {
		Response.Error(c, http.StatusBadRequest, "Invalid Request", err)
		return
	}
	book := domain.Book{
		Id:     bookModel.Id,
		Title:  bookModel.Title,
		Author: bookModel.Author,
		Year:   bookModel.Year,
		ISBN:   bookModel.ISBN,
	}

	ctx, email := h.UsecaseRequirements(c)
	if err := h.uc.CreateBook(ctx, book, email); err != nil {
		Response.Error(c, http.StatusBadRequest, "Invalid Request", err)
		return
	}

	Response.Success(c, http.StatusCreated, "Book added successfully", book)
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	ctx, email := h.UsecaseRequirements(c)
	books, err := h.uc.GetBooks(ctx, email)
	if err != nil {
		Response.Error(c, http.StatusNotFound, "Empty library", err)
		return
	}

	Response.Success(c, http.StatusOK, "Books retreived successfully", books)
}

func (h *BookHandler) GetBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		Response.Error(c, http.StatusBadRequest, "Invalid Id", err)
		return
	}
	ctx, email := h.UsecaseRequirements(c)
	book, err := h.uc.GetBook(ctx, id, email)
	if err != nil {
		Response.Error(c, http.StatusNotFound, "Id not found", err)
		return
	}
	Response.Success(c, http.StatusOK, "Book retrieved successfully", book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var bookModel dto.BookModel
	if err := c.ShouldBindJSON(&bookModel); err != nil {
		Response.Error(c, http.StatusBadRequest, "Invalid book details", err)
		return
	}

	book := domain.Book{
		Id:     bookModel.Id,
		Title:  bookModel.Title,
		Author: bookModel.Author,
		Year:   bookModel.Year,
		ISBN:   bookModel.ISBN,
	}

	ctx, email := h.UsecaseRequirements(c)
	if err := h.uc.UpdateBook(ctx, book, email); err != nil {
		Response.Error(c, http.StatusNotFound, "Id not found", err)
	}
	Response.Success(c, http.StatusOK, "Book updated successfully", book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		Response.Error(c, http.StatusBadRequest, "Invalid Id", err)
		return
	}

	ctx, email := h.UsecaseRequirements(c)
	book, err := h.uc.DeleteBook(ctx, id, email)
	if err != nil {
		Response.Error(c, http.StatusNotFound, "Id not found", err)
		return
	}
	Response.Success(c, http.StatusOK, "Book deleted successfully", book)
}
