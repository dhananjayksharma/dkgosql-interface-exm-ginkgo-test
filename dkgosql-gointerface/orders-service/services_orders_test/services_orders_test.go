package services_orders_test

import (
	"dkgosql-gointerface/orders-service/models"
	"dkgosql-gointerface/orders-service/services"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrderService", func() {
	var service *services.OrderServiceImpl

	BeforeEach(func() {
		service = services.NewOrderService()
	})

	Describe("Adding an order", func() {
		It("should add an order successfully", func() {
			order := models.Order{
				OrderID:  "1",
				Quantity: 2,
				Price:    100.0,
			}
			err := service.AddOrder(order)
			Expect(err).To(BeNil())
		})
	})

	Describe("Listing orders", func() {
		It("should list all orders", func() {
			order1 := models.Order{OrderID: "1", Quantity: 2, Price: 100.0}
			order3 := models.Order{OrderID: "6", Quantity: 2, Price: 100.0}
			order2 := models.Order{OrderID: "2", Quantity: 1, Price: 50.0}
			err := service.AddOrder(order1)
			Expect(err).To(BeNil())

			err = service.AddOrder(order2)
			Expect(err).To(BeNil())

			err = service.AddOrder(order3)
			Expect(err).To(BeNil())

			orders, err := service.ListOrders()
			Expect(err).To(BeNil())
			Expect(orders).To(HaveLen(3))
		})
	})

	Describe("Updating an order", func() {
		It("should update an existing order", func() {
			order := models.Order{OrderID: "1", Quantity: 2, Price: 100.0}
			service.AddOrder(order)

			updatedOrder := models.Order{OrderID: "1", Quantity: 5, Price: 150.0}
			err := service.UpdateOrder(updatedOrder)
			Expect(err).To(BeNil())
		})
	})

	Describe("Deleting an order", func() {
		It("should delete an order successfully", func() {
			order := models.Order{OrderID: "1", Quantity: 2, Price: 100.0}
			service.AddOrder(order)
			err := service.DeleteOrder("1")
			Expect(err).To(BeNil())

			order1 := models.Order{OrderID: "41", Quantity: 25, Price: 100.0}
			service.AddOrder(order1)
			err = service.DeleteOrder("12")
			Expect(err.Error()).To(ContainSubstring("order not exist"))
		})
	})

	Describe("IsOrderExist test", func() {
		It("should return duplicate order", func() {
			order := models.Order{
				OrderID:  "1",
				Quantity: 2,
				Price:    100.0,
			}
			err := service.AddOrder(order)
			Expect(err).To(BeNil())

			err = service.IsOrderExist(order)
			Expect(err.Error()).To(Equal(fmt.Sprintf("duplicate order id: %s", order.OrderID)))

		})
	})
})
