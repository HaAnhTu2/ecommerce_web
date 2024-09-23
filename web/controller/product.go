package controller

import (
	"net/http"
	"strconv"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"github.com/HaAnhTu2/ecommerce_web.git/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductController struct {
	ProductRepoInterface repository.ProductRepoInterface
	DB                   *mongo.Database
}

func NewProductController(ProductRepoI repository.ProductRepoInterface, db *mongo.Database) *ProductController {
	return &ProductController{ProductRepoInterface: ProductRepoI, DB: db}
}

func (p *ProductController) GetByID(c *gin.Context) {
	productid := c.Param("id")
	product, err := p.ProductRepoInterface.FindByID(c.Request.Context(), productid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
func (p *ProductController) GetAllProduct(c *gin.Context) {
	products, err := p.ProductRepoInterface.GetAll(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	// product := model.Product{
	// 	Name:        c.Request.FormValue("name"),
	// 	Description: c.Request.FormValue("description"),
	// 	Category:    c.Request.FormValue("category"),
	// }
	// stock, err := strconv.Atoi(c.Request.FormValue("stock"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock"})
	// 	return
	// }
	// product.Stock = stock

	// price, err := strconv.ParseFloat(c.Request.FormValue("price"), 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
	// 	return
	// }
	// product.Price = price
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Insert the product into the database
	products, err := p.ProductRepoInterface.Create(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": products,
	})
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	productid := c.Param("id")
	product, err := p.ProductRepoInterface.FindByID(c.Request.Context(), productid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if productname := c.PostForm("name"); productname != "" {
		product.Name = productname
	}
	if description := c.PostForm("description"); description != "" {
		product.Description = description
	}
	if prices := c.PostForm("price"); prices != "" {
		price, err := strconv.ParseFloat(prices, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid price",
			})
			return
		}
		product.Price = price
	}

	if stockStr := c.PostForm("stock"); stockStr != "" {
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid stock",
			})
			return
		}
		product.Stock = stock
	}

	if category := c.PostForm("category"); category != "" {
		product.Category = category
	}

	updatedProduct, err := p.ProductRepoInterface.Update(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update product",
			"err":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": updatedProduct,
	})
}
func (p *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid argument id",
		})
		return
	}
	if err := p.ProductRepoInterface.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
