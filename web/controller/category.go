package controller

import (
	"net/http"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"github.com/HaAnhTu2/ecommerce_web.git/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryController struct {
	CategoryRepoInterface repository.CategoryRepoInterface
	DB                    *mongo.Database
}

func NewCategoryCotroller(CategoryRepoI repository.CategoryRepoInterface, db *mongo.Database) *CategoryController {
	return &CategoryController{CategoryRepoInterface: CategoryRepoI}
}

func (ca *CategoryController) GetByID(c *gin.Context) {
	categoryid := c.Param("id")
	category, err := ca.CategoryRepoInterface.FindByID(c.Request.Context(), categoryid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}
func (ca *CategoryController) CreateCategory(c *gin.Context) {
	category := model.Category{
		Name: c.Request.FormValue("name"),
	}

	// Insert the product into the database
	category, err := ca.CategoryRepoInterface.Create(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert caregory"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func (ca *CategoryController) UpdateCategory(c *gin.Context) {
	categoryid := c.Param("id")
	category, err := ca.CategoryRepoInterface.FindByID(c.Request.Context(), categoryid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if categoryname := c.PostForm("name"); categoryname != "" {
		category.Name = categoryname
	}

	updatecategory, err := ca.CategoryRepoInterface.Update(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update category",
			"err":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": updatecategory,
	})
}
func (ca *CategoryController) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid argument id",
		})
		return
	}
	if err := ca.CategoryRepoInterface.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
