package services

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"dkgosql-gointerface/orders-service/models"
)

type OrderServiceImpl struct {
	orders []models.Order
	mu     sync.Mutex // For thread-safe access
}

// NewOrderService creates a new order service
func NewOrderService() *OrderServiceImpl {
	return &OrderServiceImpl{
		orders: make([]models.Order, 0),
	}
}

func (s *OrderServiceImpl) AddOrder(order models.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	order.Date = time.Now()
	orderExist := s.IsOrderExist(order)
	if orderExist == nil {
		s.orders = append(s.orders, order)
	} else {
		return orderExist
	}

	return nil
}

func (s *OrderServiceImpl) IsOrderExist(order models.Order) error {
	for _, row := range s.orders {
		if row.OrderID == order.OrderID {
			err := fmt.Errorf("duplicate order id: %s", order.OrderID)
			return err
		}
	}

	return nil
}

func (s *OrderServiceImpl) DeleteOrder(orderID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Printf("s.orders:%v\n\n", s.orders)
	log.Printf("s.orders:%#v\n", s.orders)
	for i, order := range s.orders {
		if order.OrderID == orderID {
			s.orders = append(s.orders[:i], s.orders[i+1:]...)
			return nil
		}
	}
	return errors.New("order not exist")
}

func (s *OrderServiceImpl) UpdateOrder(updatedOrder models.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, order := range s.orders {
		if order.OrderID == updatedOrder.OrderID {
			s.orders[i] = updatedOrder
			return nil
		}
	}
	return errors.New("order not found")
}

func (s *OrderServiceImpl) ListOrders() ([]models.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.orders, nil
}
