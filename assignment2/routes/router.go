package routes

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

// CreateRouter creates and configures a new Gin router
func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", createOrder)
	router.GET("/orders", getOrders)
	router.PUT("/orders/:orderId", updateOrder)
	router.DELETE("/orders/:orderId", deleteOrder)

	return router
}

func createOrder(c *gin.Context) {
	var newOrder Order
	if err := c.BindJSON(&newOrder); err != nil {
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
	orderID := c.Param("orderId")
	id := orderIDtoInt(orderID)
	if id == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var updatedOrder Order
	if err := c.BindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, order := range orders {
		if order.OrderID == id {
			updatedOrder.OrderID = id
			orders[i] = updatedOrder
			c.JSON(http.StatusOK, updatedOrder)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func deleteOrder(c *gin.Context) {
	orderID := c.Param("orderId")
	id := orderIDtoInt(orderID)
	if id == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	for i, order := range orders {
		if order.OrderID == id {
			orders = append(orders[:i], orders[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func orderIDtoInt(orderID string) int {
	id, err := strconv.Atoi(orderID)
	if err != nil {
		return -1
	}
	return id
}
