package main

import (
	"log"

	"github.com/bagashiz/pustaka-api/book"
	"github.com/bagashiz/pustaka-api/handler"
	"github.com/bagashiz/pustaka-api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepo := book.NewRepository(db)
	bookService := book.NewService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)

	v1.POST("/books", bookHandler.CreateBook)

	v1.PUT("/books/:id", bookHandler.UpdateBook)

	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(config.HTTPServerAddress)
}
