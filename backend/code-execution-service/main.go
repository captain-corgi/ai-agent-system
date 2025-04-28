// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	deliveryhttp "github.com/captain-corgi/ai-agent-system/code-execution-service/delivery/http"
)

func main() {
	router := gin.Default()
	deliveryhttp.RegisterRoutes(router)
	log.Println("[code-execution-service] Starting server on :8082")
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
