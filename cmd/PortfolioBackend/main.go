package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/PortfolioBackend/internal/routes"
	"log"
)

func main() {
	gin.ForceConsoleColor()

	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1", "::1"}) // For IPv4 and IPv6 localhost
	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}
	routes.Routes(router)

	router.Run("localhost:8080")
}
