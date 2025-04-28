// infrastructure/ollama_client.go
package infrastructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/captain-corgi/ai-agent-system/ai-service/domain"
)

// GeneratePlanWithOllama calls the Ollama LLM REST API and returns a plan.
func GeneratePlanWithOllama(taskType string, payload map[string]interface{}) (*domain.Plan, error) {
	ollamaURL := os.Getenv("OLLAMA_API_URL")
	if ollamaURL == "" {
		ollamaURL = "http://localhost:11434/api/plan" // Default endpoint
	}

	reqBody := map[string]interface{}{
		"task_type": taskType,
		"payload": payload,
	}
	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", ollamaURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama returned status %d", resp.StatusCode)
	}

	var plan domain.Plan
	if err := json.NewDecoder(resp.Body).Decode(&plan); err != nil {
		return nil, errors.New("invalid response from ollama")
	}
	if plan.ID == "" || len(plan.Steps) == 0 {
		return nil, errors.New("ollama returned incomplete plan")
	}
	return &plan, nil
}
