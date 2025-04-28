// delivery/http/handler.go
package http

import (
	"net/http"
)

import (
	"github.com/gin-gonic/gin"
	"github.com/captain-corgi/ai-agent-system/ai-service/usecase"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/ai/plan", generatePlanHandler)
	router.GET("/ai/health", healthHandler)
}

func generatePlanHandler(c *gin.Context) {
	var req struct {
		TaskType string                 `json:"task_type" binding:"required,oneof=code browse fs"`
		Payload  map[string]interface{} `json:"payload" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plan, err := usecase.GeneratePlan(req.TaskType, req.Payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate plan"})
		return
	}
	c.JSON(http.StatusOK, plan)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

