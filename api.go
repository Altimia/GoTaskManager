package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitAPI() {
	router = gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

func CloseAPI() {
	// code for closing the API
	// This function will be implemented later when we have a way to gracefully shutdown the server
}
