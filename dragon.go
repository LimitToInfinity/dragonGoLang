package main

import "github.com/gin-gonic/gin"

// structs are objects/classes
type Dragon struct {
	Name string;
	Breed string;
	Age int;
	Female bool;
}

func handleDragon(context *gin.Context) {
	dragon := getDragon("Roar")
	context.JSON(200, dragon)
}

func getDragon(name string) Dragon {
	return Dragon{
		Name: name,
		Breed: "destroyer",
		Age: 3727,
		Female: true,
	}
}