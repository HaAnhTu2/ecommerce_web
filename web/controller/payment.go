package controller

import (
	"net/http"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"github.com/HaAnhTu2/ecommerce_web.git/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentController struct {
	PaymentRepoInterface repository.PaymentRepoInterface
	DB                   *mongo.Database
}

func NewPaymentController(PaymentRepoI repository.PaymentRepoInterface, db *mongo.Database) *PaymentController {
	return &PaymentController{PaymentRepoInterface: PaymentRepoI, DB: db}
}

func (pm *PaymentController) CreatePayment(c *gin.Context) {
	payment := model.Payment{
		PaymentMethod: c.Request.FormValue("payment_method"),
		Status:        c.Request.FormValue("status"),
	}
	// Insert the payment into the database
	payment, err := pm.PaymentRepoInterface.Create(c.Request.Context(), payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert payment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"payment": payment,
	})
}
func (pm *PaymentController) UpdatePayment(c *gin.Context) {
	paymentid := c.Param("id")
	payment, err := pm.PaymentRepoInterface.FindByID(c.Request.Context(), paymentid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	if paymentmethod := c.PostForm("payment_method"); paymentmethod != "" {
		payment.PaymentMethod = paymentmethod
	}
	if status := c.PostForm("status"); status != "" {
		payment.Status = status
	}
	updatedPayment, err := pm.PaymentRepoInterface.Update(c.Request.Context(), payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update payment",
			"err":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payment": updatedPayment,
	})
}
func (pm *PaymentController) DeletePayment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	if err := pm.PaymentRepoInterface.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
