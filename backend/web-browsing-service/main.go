// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	deliveryhttp "github.com/captain-corgi/ai-agent-system/web-browsing-service/delivery/http"
)

func main() {
	router := gin.Default()
	deliveryhttp.RegisterRoutes(router)
	log.Println("[web-browsing-service] Starting server on :8084")
	if err := router.Run(":8084"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
