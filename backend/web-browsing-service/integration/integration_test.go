package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	deliveryhttp "github.com/captain-corgi/ai-agent-system/web-browsing-service/delivery/http"
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
	req, _ := http.NewRequest("GET", "/browse/health", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestBrowseURLEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body := map[string]interface{}{"url": "https://example.com"}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/browse/url", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Browsed https://example.com")
}

func TestBrowseURLEndpoint_InvalidInput(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body := map[string]interface{}{"url": "not-a-url"}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/browse/url", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
