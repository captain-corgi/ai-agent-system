// delivery/http/handler.go
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/captain-corgi/ai-agent-system/web-browsing-service/usecase"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/browse/url", browseURLHandler)
	router.GET("/browse/health", healthHandler)
}

func browseURLHandler(c *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required,url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := usecase.BrowseURL(req.URL)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
