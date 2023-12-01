package main

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	server *http.Server
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}

func InitAPI() {
	router = setupRouter()
	server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("ListenAndServe failed", zap.Error(err))
		}
	}()
}

func CloseAPI() {
	// Use the global server variable for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Error("Error shutting down the server", zap.Error(err))
	} else {
		zap.L().Info("Server shut down gracefully")
	}
}
