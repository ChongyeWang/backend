package main

import (
	database "backend/db"
	"backend/handlers"
	"backend/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := database.InitDB(); err != nil {
		panic(err)
	}

	// Create Gin router
	r := gin.Default()

	// REST API routes
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)

	// WebSocket route
	r.GET("/ws", websocket.HandleWebSocket)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
