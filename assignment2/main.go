package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Order represents the structure of an order
type Order struct {
	OrderID      int       `json:"orderId"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items"`
}

// Item represents the structure of an item in an order
type Item struct {
	LineItemID  int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

var orders []Order

func main() {
	r := gin.Default()

	r.POST("/orders", createOrder)
	r.GET("/orders", getOrders)
	r.PUT("/orders/:orderId", updateOrder)
	r.DELETE("/orders/:orderId", deleteOrder)

	r.Run(":8080")
}

func createOrder(c *gin.Context) {
	var newOrder Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder.OrderID = len(orders) + 1
	orders = append(orders, newOrder)

	c.JSON(http.StatusCreated, newOrder)
}

func getOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}

func updateOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var updatedOrder Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, order := range orders {
		if order.OrderID == orderID {
			updatedOrder.OrderID = orderID
			orders[i] = updatedOrder
			c.JSON(http.StatusOK, updatedOrder)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func deleteOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	for i, order := range orders {
		if order.OrderID == orderID {
			orders = append(orders[:i], orders[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}
