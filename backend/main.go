package main

// Author : Antony Injila
// Contributors : Nil :)
// Project name : DrivePro
// Github : https://github.com/AntonyIS/DrivePro
// Description : Platform for buying and selling second hand cars
// Technologies : Golang,Go Gin (Go REST API framework) Docker, AWS

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin router to router requests to the server
	router := gin.Default()

	// Home route
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"res": "Hello DrivePro"})
	})

	router.Run(":5000")
}
