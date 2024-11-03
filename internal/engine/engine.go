package engine

import (
	"oaiproxy/internal/handlers"
	"oaiproxy/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Run() {

	// Create an engine instance
	engine := gin.Default()

	// Load all HTML templates
	engine.LoadHTMLGlob("templates/*")

	// Register handlers

	// Welcome page
	engine.GET("/", handlers.WelcomeHandler)

	// Create a router group
	engine.Group("/")

	// Use the API key middleware for validation
	engine.Use(middlewares.ApiKeyMiddleware())

	// Handle all the other routes
	engine.NoRoute(handlers.ProxyHandler)

	// Run the server
	engine.Run(":8080")
}
