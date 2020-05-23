package api

import (
	_ "github.com/MarianoArias/challange-api/cmd/items/docs"
	"github.com/MarianoArias/challange-api/internal/app/items/controller"
	"github.com/MarianoArias/challange-api/internal/pkg/health-handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/zsais/go-gin-prometheus"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/images/:fileName", controller.GetImageHandler)

	router.GET("/items/", controller.CgetHandler)
	router.DELETE("/items/:id", controller.DeleteHandler)
	router.GET("/items/:id", controller.GetHandler)
	router.PATCH("/items/:id", controller.PatchHandler)
	router.POST("/items/", controller.PostHandler)

	// Api Doc Endpoint => /doc/index.html
	router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Health Endpoint => /health/
	router.GET("/health/", healthhandler.HealthHandler)
	// Metrics Endpoint => /metrics/
	prometheus := ginprometheus.NewPrometheus("gin")
	prometheus.Use(router)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
