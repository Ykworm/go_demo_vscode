package routes

import (
	"gin-demo-project/internal/handlers" // Ensure this package is installed
	"math/rand"                          // Ensure this package is installed
	"time"                               // Ensure this package is installed

	"github.com/gin-gonic/gin" // Ensure this package is installed
)

func SetupRoutes(router *gin.Engine, handler *handlers.Handler) {
	// router.GET("/", handler.GetHome)
	router.GET("/test", testHandler)
}

func testHandler(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 2000
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	c.String(200, string(result))
}
