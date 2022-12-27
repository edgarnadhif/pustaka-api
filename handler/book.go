package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bagashiz/pustaka-api/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// bookHandler implements book.Service struct
type bookHandler struct {
	bookService book.Service
}

// NewBookHandler creates a new bookHandler instance
func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

// GetBook retrieves a single data from the Books table based on the requested ID
func (h *bookHandler) GetBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// GetBooks retrieves all data from the Books table
func (h *bookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

// CreateBook creates a new data to the Books table based on requested parameters
func (h *bookHandler) CreateBook(ctx *gin.Context) {
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

	bookResponse := convertToBookResponse(book)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// UpdateBook updates existing data from the Books table based on the requested parameters
func (h *bookHandler) UpdateBook(ctx *gin.Context) {
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

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(book)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// DeleteBook removes data from the Books table based on requested id
func (h *bookHandler) DeleteBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// convertToBookResponse converts data from book.BookRequest struct to book.BookResponse struct
func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Rating:      b.Rating,
		Price:       b.Price,
		Discount:    b.Discount,
	}
}
