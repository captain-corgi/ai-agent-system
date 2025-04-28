// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	deliveryhttp "github.com/captain-corgi/ai-agent-system/filesystem-service/delivery/http"
)

func main() {
	router := gin.Default()
	deliveryhttp.RegisterRoutes(router)
	log.Println("[filesystem-service] Starting server on :8083")
	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
