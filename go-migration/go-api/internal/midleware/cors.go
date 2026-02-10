package midleware

import "github.com/gin-contrib/cors"

func CorsConfig() cors.Config {
	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Email", "User-Agent-String"},
		AllowCredentials: true,
	}
	return config
}
