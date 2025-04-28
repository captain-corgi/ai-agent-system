package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	deliveryhttp "github.com/captain-corgi/ai-agent-system/code-execution-service/delivery/http"
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
	req, _ := http.NewRequest("GET", "/execute/health", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestExecuteCodeEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body := map[string]interface{}{
		"language": "python",
		"source": "print('hello')",
	}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/execute/code", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Executed python code!")
}

func TestExecuteCodeEndpoint_InvalidInput(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body := map[string]interface{}{
		"language": "",
		"source": "",
	}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/execute/code", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
