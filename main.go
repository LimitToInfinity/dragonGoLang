package main

import "github.com/gin-gonic/gin"

func main() {
	// capital D on Default means exported (public method)
	// := is go syntax for creating new variable inside a func
	// go likes single letter variables because
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}