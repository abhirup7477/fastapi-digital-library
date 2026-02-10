package midleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/abhirup7477/go-books-api/internal/delivery/http/Response"
	"github.com/gin-gonic/gin"
)

func LoggerMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_agent_string := c.Request.Header.Get("User-Agent-String")
		if user_agent_string == "" {
			Response.Error(c, http.StatusUnauthorized, "No user agent found in the header", errors.New("User-Agent-String missing"))
			return
		}
		log.Println("Request Received from:", user_agent_string)

		c.Next()
	}
}
