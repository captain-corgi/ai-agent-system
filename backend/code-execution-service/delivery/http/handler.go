// delivery/http/handler.go
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/captain-corgi/ai-agent-system/code-execution-service/usecase"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/execute/code", executeCodeHandler)
	router.GET("/execute/health", healthHandler)
}

func executeCodeHandler(c *gin.Context) {
	var req struct {
		Language string `json:"language" binding:"required"`
		Source   string `json:"source" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := usecase.ExecuteCode(req.Language, req.Source)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
