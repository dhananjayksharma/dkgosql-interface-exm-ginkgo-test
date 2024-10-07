package models

import "time"

// Order represents an order with basic details
type Order struct {
	OrderID  string    `json:"order_id"`
	Quantity int       `json:"quantity"`
	Price    float64   `json:"price"`
	Date     time.Time `json:"date"`
}

// OrderService defines the operations to manage orders
type OrderService interface {
	AddOrder(order Order) error
	IsOrderExist(order Order) error
	DeleteOrder(orderID string) error
	UpdateOrder(order Order) error
	ListOrders() ([]Order, error)
}
