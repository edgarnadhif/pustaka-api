package handler

import (
	"fmt"
	"net/http"

	"github.com/bagashiz/pustaka-api/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Bagas Hizbullah",
		"bio":  "Information Systems student",
	})
}

func (h *bookHandler) HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title": "hello world!",
	})
}

func (h *bookHandler) BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")

	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *bookHandler) QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func (h *bookHandler) PostBooksHandler(ctx *gin.Context) {
	var bookRequest book.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %v, condition: %v", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
