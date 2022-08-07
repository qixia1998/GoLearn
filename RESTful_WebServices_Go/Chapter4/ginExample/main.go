package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("pingTime", func(c *gin.Context) {
		// JSON serializer is available on gin context
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	r.Run(":8000") // Listen and serve on 0.0.0.0:8000
}
