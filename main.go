package main

import (
	"github.com/cbonoz/http-proxy/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(logger.SetLogger())

	r.POST("/proxy", controllers.ProxyRequest)
	r.GET("up", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":8080")
}
