package controller

import (
	"net/http"
	"strconv"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"github.com/HaAnhTu2/ecommerce_web.git/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemController struct {
	OrderItemInterface repository.OrderItemInterface
	DB                 *mongo.Database
}

func NewOrderItemController(OrderItemI repository.OrderItemInterface, db *mongo.Database) *OrderItemController {
	return &OrderItemController{OrderItemInterface: OrderItemI, DB: db}
}
func (oi *OrderItemController) GetByID(c *gin.Context) {
	order_item_id := c.Param("id")
	orderitem, err := oi.OrderItemInterface.FindByID(c.Request.Context(), order_item_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"orderitem": orderitem,
	})
}

func (oi *OrderItemController) CreateOrderItem(c *gin.Context) {
	var orderitem model.OrderItem
	quantity, err := strconv.Atoi(c.Request.FormValue("quantity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}
	orderitem.Quantity = quantity

	price, err := strconv.ParseFloat(c.Request.FormValue("price"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
		return
	}
	orderitem.Price = price
	// Insert the order into the database
	orderitems, err := oi.OrderItemInterface.Create(c.Request.Context(), orderitem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert orderitem"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orderitem": orderitems,
	})
}
func (oi *OrderItemController) UpdateOrderItem(c *gin.Context) {
	orderitemid := c.Param("id")
	orderitem, err := oi.OrderItemInterface.FindByID(c.Request.Context(), orderitemid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order Item not found"})
		return
	}
	if quantityStr := c.PostForm("quantity"); quantityStr != "" {
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid quantity",
			})
			return
		}
		orderitem.Quantity = quantity
	}
	if prices := c.PostForm("price"); prices != "" {
		price, err := strconv.ParseFloat(prices, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid price",
			})
			return
		}
		orderitem.Price = price
	}

	updatedOrderItem, err := oi.OrderItemInterface.Update(c.Request.Context(), orderitem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update order item",
			"err":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_item": updatedOrderItem,
	})
}
func (oi *OrderItemController) DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	if err := oi.OrderItemInterface.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
