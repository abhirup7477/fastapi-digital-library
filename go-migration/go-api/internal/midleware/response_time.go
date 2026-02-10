package midleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func ResponseTimeMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		responseTime := time.Since(start)
		log.Printf("%s %s Response-Time: %v", c.Request.Method, c.Request.URL.Path, responseTime)
		c.Writer.Header().Set("Response-Time", responseTime.String())
	}
}
