package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	deliveryhttp "github.com/captain-corgi/ai-agent-system/filesystem-service/delivery/http"
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
	req, _ := http.NewRequest("GET", "/fs/health", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestReadFileEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body := map[string]interface{}{"path": "/tmp/test.txt"}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/fs/read", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Read from /tmp/test.txt")
}

func TestWriteFileEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body := map[string]interface{}{"path": "/tmp/test.txt", "content": "hello"}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/fs/write", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Wrote to /tmp/test.txt")
}

func TestReadFileEndpoint_InvalidInput(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body := map[string]interface{}{"path": ""}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/fs/read", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
