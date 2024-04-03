package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	productHandler "github.com/products-api/products/infraesturucture/handler"
)

type Handlers struct {
	productsHandler productHandler.IProductHandler
}

func NewHandlers() *Handlers {
	handlers := &Handlers{
		productsHandler: productHandler.BuildProductsHandler(),
	}
	return handlers
}

func (h Handlers) LoadRoutes(r *gin.Engine) *gin.Engine {
	r.GET("", func(c *gin.Context) {
		resp := struct {
			Status  string    `json:"status"`
			Date    time.Time `json:"date"`
			Message string    `json:"message"`
		}{
			Status:  "Ok",
			Date:    time.Now(),
			Message: "Server is running",
		}
		c.JSON(http.StatusOK, gin.H{"body": resp})
	})

	v1 := r.Group("v1", func(ctx *gin.Context) {})
	{
		v1.GET("products", h.productsHandler.Get)
		v1.POST("products", h.productsHandler.Post)
	}
	return r
}
