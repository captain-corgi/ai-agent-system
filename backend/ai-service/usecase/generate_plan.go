// usecase/generate_plan.go
package usecase

import (
	"errors"
	"github.com/captain-corgi/ai-agent-system/ai-service/domain"
	"github.com/captain-corgi/ai-agent-system/ai-service/infrastructure"
)

func GeneratePlan(taskType string, payload map[string]interface{}) (*domain.Plan, error) {
	if taskType != "code" && taskType != "browse" && taskType != "fs" {
		return nil, errors.New("invalid task type")
	}
	plan, err := infrastructure.GeneratePlanWithOllama(taskType, payload)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

