package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/products-api/products/application/usecases"
)

type IProductHandler interface {
	Get(c *gin.Context)
}

type ProductHandler struct {
	getAllProductsUC usecases.IGetAllProducts
}

func NewProductsHandler(getAllProductsUC usecases.IGetAllProducts) IProductHandler {
	return &ProductHandler{
		getAllProductsUC: getAllProductsUC,
	}
}

func (ph ProductHandler) Get(c *gin.Context) {
	result, err := ph.getAllProductsUC.Execute()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"body": result})
}
