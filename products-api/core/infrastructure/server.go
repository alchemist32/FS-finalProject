package infrastructure

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/products-api/core/infrastructure/router"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.RedirectTrailingSlash = true

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"HEAD", "DELETE", "PUT", "POST", "GET", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Okta-Id", "Content-Security-Policy"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "X-version", "X-commit-hash", "X-build"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           1 * time.Hour,
	}))

	// fmt.Println("I'm a Change")
	handlers := router.NewHandlers()
	handlers.LoadRoutes(r)

	return r
}
