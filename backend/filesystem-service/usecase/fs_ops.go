package usecase

import (
	"errors"
	"fmt"
	"github.com/captain-corgi/ai-agent-system/filesystem-service/domain"
)

// ReadFile mocks reading a file. Replace with real FS logic.
func ReadFile(path string) (*domain.FsJob, error) {
	if path == "" {
		return nil, errors.New("path is required")
	}
	job := &domain.FsJob{
		ID:     "fsjob-1",
		Path:   path,
		OpType: "read",
		Status: "completed",
		Result: fmt.Sprintf("Read from %s", path),
	}
	return job, nil
}

// WriteFile mocks writing to a file. Replace with real FS logic.
func WriteFile(path, content string) (*domain.FsJob, error) {
	if path == "" || content == "" {
		return nil, errors.New("path and content are required")
	}
	job := &domain.FsJob{
		ID:      "fsjob-2",
		Path:    path,
		OpType:  "write",
		Content: content,
		Status:  "completed",
		Result:  fmt.Sprintf("Wrote to %s", path),
	}
	return job, nil
}
