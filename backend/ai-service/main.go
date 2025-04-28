// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/captain-corgi/ai-agent-system/ai-service/delivery/http"
)

func main() {
	router := gin.Default()
	http.RegisterRoutes(router)
	log.Println("[ai-service] Starting server on :8081")
	err := router.Run(":8081")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
