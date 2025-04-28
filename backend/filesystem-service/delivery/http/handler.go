// delivery/http/handler.go
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/captain-corgi/ai-agent-system/filesystem-service/usecase"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/fs/read", readFileHandler)
	router.POST("/fs/write", writeFileHandler)
	router.GET("/fs/health", healthHandler)
}

func readFileHandler(c *gin.Context) {
	var req struct {
		Path string `json:"path" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := usecase.ReadFile(req.Path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func writeFileHandler(c *gin.Context) {
	var req struct {
		Path    string `json:"path" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := usecase.WriteFile(req.Path, req.Content)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
