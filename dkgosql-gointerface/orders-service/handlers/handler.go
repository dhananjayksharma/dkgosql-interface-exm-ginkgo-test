package handlers

import (
	"net/http"

	"dkgosql-gointerface/orders-service/models"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Service models.OrderService
}

// NewOrderHandler creates a new handler for managing orders
func NewOrderHandler(service models.OrderService) *OrderHandler {
	return &OrderHandler{
		Service: service,
	}
}

// AddOrder handler
func (h *OrderHandler) AddOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.AddOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order added successfully!"})
}

// DeleteOrder handler
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	orderID := c.Param("order_id")
	if err := h.Service.DeleteOrder(orderID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully!"})
}

// UpdateOrder handler
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.UpdateOrder(order); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully!"})
}

// ListOrders handler
func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.Service.ListOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
