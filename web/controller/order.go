package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"github.com/HaAnhTu2/ecommerce_web.git/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderController struct {
	OrderRepoInterface repository.OrderRepoInterface
	DB                 *mongo.Database
}

func NewOrderController(OrderRepoI repository.OrderRepoInterface, db *mongo.Database) *OrderController {
	return &OrderController{OrderRepoInterface: OrderRepoI, DB: db}
}
func (o *OrderController) GetByID(c *gin.Context) {
	orderid := c.Param("id")
	order, err := o.OrderRepoInterface.FindByID(c.Request.Context(), orderid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	order := model.Order{
		Status: c.Request.FormValue("status"),
	}
	order.OrderDate = time.Now()
	// Insert the order into the database
	orders, err := o.OrderRepoInterface.Create(c.Request.Context(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order": orders,
	})
}
func (o *OrderController) UpdateOrder(c *gin.Context) {
	orderid := c.Param("id")
	order, err := o.OrderRepoInterface.FindByID(c.Request.Context(), orderid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if totalamountStr := c.PostForm("total_amount"); totalamountStr != "" {
		totalamount, err := strconv.ParseFloat(totalamountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid total amount",
			})
			return
		}
		order.TotalAmount = totalamount
	}
	if status := c.PostForm("status"); status != "" {
		order.Status = status
	}

	order.OrderDate = time.Now()
	updatedOrder, err := o.OrderRepoInterface.Update(c.Request.Context(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update order",
			"err":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order": updatedOrder,
	})
}
func (o *OrderController) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	if err := o.OrderRepoInterface.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
