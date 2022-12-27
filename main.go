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

	//* CRUD
	var books []book.Book
	var book book.Book

	// CREATE
	book.Title = "Clean Code"
	book.Price = 40
	book.Discount = 10
	book.Rating = 5
	book.Description = "The best programming book."

	err = db.Debug().Create(&book).Error
	if err != nil {
		fmt.Printf("error creating book record: %v\n", err)
	}

	// READ
	err = db.Debug().First(&book).Error
	if err != nil {
		fmt.Printf("error finding book record: %v\n", err)
	}

	fmt.Printf("Title\t: %v\nBook object: %v\n", book.Title, book)

	err = db.Debug().Where("rating = ?", 5).Find(&books).Error
	if err != nil {
		fmt.Printf("error finding books record: %v\n", err)
	}

	for _, b := range books {
		fmt.Printf("Title\t: %v\nBook object: %v\n", b.Title, b)
	}

	// UPDATE
	book.Title = "Clean Coder"
	book.Price = 30
	book.Discount = 15
	book.Rating = 5
	book.Description = "The 2nd best programming book."

	err = db.Debug().Save(&book).Error
	if err != nil {
		fmt.Printf("error updating book record: %v\n", err)
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
