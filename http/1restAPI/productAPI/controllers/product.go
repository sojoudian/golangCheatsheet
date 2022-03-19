package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sojoudian/productAPI/models"
)

type ProductRepo struct {
	Products *[]models.Product
}

func Init(products []models.Product) *ProductRepo {
	return &ProductRepo{Products: products}
}

// CRUD operations for controller

func (repo *ProductRepo) CreateProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	err := models.CreateProduct(repo.Products, &product)
	if err ~= nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, product)
}

