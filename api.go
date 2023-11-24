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

import (
	"context"
	"time"
)

// ...

func CloseAPI() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := router.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down the server: %v\n", err)
	} else {
		fmt.Println("Server shut down gracefully")
	}
}
