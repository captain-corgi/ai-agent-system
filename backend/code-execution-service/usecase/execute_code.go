package usecase

import (
	"errors"
	"fmt"
	"github.com/captain-corgi/ai-agent-system/code-execution-service/domain"
)

// ExecuteCode mocks code execution. Replace with Docker integration for real runs.
func ExecuteCode(language, source string) (*domain.CodeJob, error) {
	if language == "" || source == "" {
		return nil, errors.New("language and source are required")
	}
	// TODO: Integrate with Docker for real code execution
	job := &domain.CodeJob{
		ID:     "job-123",
		Lang:   language,
		Code:   source,
		Status: "completed",
		Output: fmt.Sprintf("Executed %s code!", language),
	}
	return job, nil
}
