package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	deliveryhttp "github.com/captain-corgi/ai-agent-system/ai-service/delivery/http"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	deliveryhttp.RegisterRoutes(r)
	return r
}

func TestHealthEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ai/health", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestGeneratePlanEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body := map[string]interface{}{
		"task_type": "code",
		"payload": map[string]interface{}{"foo": "bar"},
	}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/ai/plan", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatalf("Expected 200 OK but got %d. Response: %s. Is Ollama running?", w.Code, w.Body.String())
	}
	var resp struct {
		ID    string      `json:"id"`
		Steps interface{} `json:"steps"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Response is not valid JSON: %v", err)
	}
	if resp.ID == "" {
		t.Errorf("plan ID is empty in response: %s", w.Body.String())
	}
	if resp.Steps == nil {
		t.Errorf("steps missing in response: %s", w.Body.String())
	}
}

func TestGeneratePlanEndpoint_InvalidType(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body := map[string]interface{}{
		"task_type": "invalid",
		"payload": map[string]interface{}{"foo": "bar"},
	}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/ai/plan", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
