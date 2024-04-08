package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/products-api/products/application/usecases"
	"github.com/products-api/products/domain/models"
	"github.com/products-api/products/infraesturucture/dao"
)

type IProductHandler interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	GetByBarcode(c *gin.Context)
}

type ProductHandler struct {
	getAllProductsUC    usecases.IGetAllProducts
	addProductsUC       usecases.IAddProducts
	getProductByBarcode usecases.IGetProductByBarcode
}

func NewProductsHandler(
	getAllProductsUC usecases.IGetAllProducts,
	addProductsUC usecases.IAddProducts,
	getByBarcode usecases.IGetProductByBarcode) IProductHandler {
	return &ProductHandler{
		getAllProductsUC:    getAllProductsUC,
		addProductsUC:       addProductsUC,
		getProductByBarcode: getByBarcode,
	}
}

func (ph ProductHandler) Get(c *gin.Context) {
	result, err := ph.getAllProductsUC.Execute()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"body": result})
}

func (ph ProductHandler) Post(c *gin.Context) {
	var productInput models.Product
	errMarshal := c.ShouldBindJSON(&productInput)

	if errMarshal != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error reading payload"})
		return
	}
	addErr := ph.addProductsUC.Execute(productInput)

	if addErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "cannot add product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})
}

func (ph ProductHandler) GetByBarcode(c *gin.Context) {
	barcode, _ := c.Params.Get("barcode")
	product, err := ph.getProductByBarcode.Execute(barcode)

	if err != nil && errors.Is(dao.NotFoundProduct, err) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "The product with that barcode does not exist"})
		return
	}

	if err != nil {
		fmt.Printf("Error trying to get the product: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, product)
}
