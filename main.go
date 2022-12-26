package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "Bagas Hizbullah",
			"bio":  "Information Systems student",
		})
	})

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"title": "Hello World!",
		})
	})

	router.Run(":8888")
}
