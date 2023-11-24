package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

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

	// Define the server with a specific address and attach the router
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Run the server in a goroutine so that it doesn't block
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe: %v\n", err)
		}
	}()
}

// Define a global server variable to be used for graceful shutdown
var server *http.Server

func InitAPI() {
	router = gin.Default()

	// ... (other code remains unchanged)

	// Initialize the server with the router
	server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// ... (other code remains unchanged)
}

func CloseAPI() {
	// Use the global server variable for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down the server: %v\n", err)
	} else {
		fmt.Println("Server shut down gracefully")
	}
}
