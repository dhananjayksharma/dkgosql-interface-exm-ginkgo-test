package main

import (
	"dkgosql-gointerface/orders-service/handlers"
	"dkgosql-gointerface/orders-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize service and handler
	orderService := services.NewOrderService()
	orderHandler := handlers.NewOrderHandler(orderService)

	// Set up routes
	r.POST("/orders", orderHandler.AddOrder)
	r.DELETE("/orders/:order_id", orderHandler.DeleteOrder)
	r.PUT("/orders", orderHandler.UpdateOrder)
	r.GET("/orders", orderHandler.ListOrders)

	// Run the server
	r.Run(":8119")
}
