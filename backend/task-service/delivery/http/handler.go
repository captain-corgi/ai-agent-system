// delivery/http/handler.go
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/captain-corgi/ai-agent-system/task-service/infrastructure"
	"github.com/captain-corgi/ai-agent-system/task-service/usecase"
)

type Handler struct {
	repo infrastructure.TaskRepository
}

func NewHandler(repo infrastructure.TaskRepository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.POST("/tasks", h.createTaskHandler)
	router.GET("/tasks", h.listTasksHandler)
	router.GET("/tasks/:id", h.getTaskHandler)
	router.GET("/tasks/health", healthHandler)
}

func (h *Handler) createTaskHandler(c *gin.Context) {
	var req struct {
		Type    string      `json:"type" binding:"required,oneof=code browse fs"`
		Payload interface{} `json:"payload" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := usecase.CreateTask(c.Request.Context(), h.repo, req.Type, req.Payload)
	if err != nil {
		if err == usecase.ErrInvalidTaskType {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *Handler) listTasksHandler(c *gin.Context) {
	tasks, err := usecase.ListTasks(c.Request.Context(), h.repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) getTaskHandler(c *gin.Context) {
	id := c.Param("id")
	task, err := usecase.GetTask(c.Request.Context(), h.repo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get task"})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
