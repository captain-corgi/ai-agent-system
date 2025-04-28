// usecase/create_task.go
package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/captain-corgi/ai-agent-system/task-service/domain"
	"github.com/captain-corgi/ai-agent-system/task-service/infrastructure"
)

var ErrInvalidTaskType = errors.New("invalid task type")

func CreateTask(ctx context.Context, repo infrastructure.TaskRepository, taskType string, payload interface{}) (*domain.Task, error) {
	if taskType != "code" && taskType != "browse" && taskType != "fs" {
		return nil, ErrInvalidTaskType
	}
	id := fmt.Sprintf("task-%d", time.Now().UnixNano())
	now := time.Now().UTC().Format(time.RFC3339)
	task := &domain.Task{
		ID:        id,
		Type:      taskType,
		Status:    "pending",
		Payload:   fmt.Sprintf("%v", payload),
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := repo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func ListTasks(ctx context.Context, repo infrastructure.TaskRepository) ([]*domain.Task, error) {
	return repo.ListAll(ctx)
}

func GetTask(ctx context.Context, repo infrastructure.TaskRepository, id string) (*domain.Task, error) {
	return repo.GetByID(ctx, id)
}
