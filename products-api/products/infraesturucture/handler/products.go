package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/products-api/products/application/usecases"
	"github.com/products-api/products/domain/models"
)

type IProductHandler interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
}

type ProductHandler struct {
	getAllProductsUC usecases.IGetAllProducts
	addProductsUC    usecases.IAddProducts
}

func NewProductsHandler(getAllProductsUC usecases.IGetAllProducts, addProductsUC usecases.IAddProducts) IProductHandler {
	return &ProductHandler{
		getAllProductsUC: getAllProductsUC,
		addProductsUC:    addProductsUC,
	}
}

func (ph ProductHandler) Get(c *gin.Context) {
	result, err := ph.getAllProductsUC.Execute()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"body": result})
}

func (ph ProductHandler) Post(c *gin.Context) {
	var productInput models.Product
	errMarshal := c.ShouldBindJSON(&productInput)

	if errMarshal != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error reading payload"})
	}
	addErr := ph.addProductsUC.Execute(productInput)

	if addErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "cannot add product"})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})

}
