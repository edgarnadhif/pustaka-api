package main

import (
	"fmt"
	"log"

	"github.com/bagashiz/pustaka-api/book"
	"github.com/bagashiz/pustaka-api/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=password dbname=pustaka_api port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepo := book.NewRepository(db)

	book := book.Book{
		Title:       "The Pragmatic Engineer",
		Description: "The 3rd best programming book.",
		Price:       50,
		Discount:    0,
		Rating:      5,
	}

	_, err = bookRepo.Create(book)
	if err != nil {
		fmt.Printf("error creating data: %v\n", err)
	}

	book, err = bookRepo.FindByID(14)
	if err != nil {
		fmt.Printf("error retrieving data: %v\n", err)
	}

	fmt.Printf("Title\t: %v\n", book.Title)

	books, err := bookRepo.FindAll()
	if err != nil {
		fmt.Printf("error retrieving data: %v\n", err)
	}

	for _, book := range books {
		fmt.Printf("Title\t: %v\n", book.Title)
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)

	router.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
