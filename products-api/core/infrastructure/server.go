package infrastructure

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/products-api/core/infrastructure/adapter/client/mock"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.RedirectTrailingSlash = true
	mock.NewMockClient()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"HEAD", "DELETE", "PUT", "POST", "GET", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Okta-Id", "Content-Security-Policy"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "X-version", "X-commit-hash", "X-build"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           1 * time.Hour,
	}))

	return r
}
