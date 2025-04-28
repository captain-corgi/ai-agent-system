package integration

import (
	"bytes"
	"encoding/json"
	"os"
	"net/http"
	"net/http/httptest"
	"testing"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/captain-corgi/ai-agent-system/task-service/infrastructure"
	handler "github.com/captain-corgi/ai-agent-system/task-service/delivery/http"
)

func setupTestRouter(t *testing.T) *gin.Engine {
	os.Remove("test_task_service.db")
	db, err := sql.Open("sqlite3", "test_task_service.db")
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	t.Cleanup(func() { db.Close(); os.Remove("test_task_service.db") })
	if err := infrastructure.Migrate(db); err != nil {
		t.Fatalf("failed to migrate test db: %v", err)
	}
	repo := infrastructure.NewSQLiteTaskRepository(db)
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	h := handler.NewHandler(repo)
	h.RegisterRoutes(r)

	return r
}

func TestHealthEndpoint(t *testing.T) {
	r := setupTestRouter(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/health", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestCreateAndGetTask(t *testing.T) {
	r := setupTestRouter(t)
	w := httptest.NewRecorder()
	body := map[string]interface{}{ "type": "code", "payload": map[string]interface{}{"foo": "bar"} }
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var created map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &created)
	id := created["id"].(string)

	// Get by ID
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/tasks/"+id, nil)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)
	assert.Contains(t, w2.Body.String(), id)
}

func TestListTasks(t *testing.T) {
	r := setupTestRouter(t)
	// Create two tasks
	for i := 0; i < 2; i++ {
		w := httptest.NewRecorder()
		body := map[string]interface{}{ "type": "code", "payload": map[string]interface{}{"foo": i} }
		jsonBytes, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	}
	// List
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var tasks []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &tasks)
	assert.GreaterOrEqual(t, len(tasks), 2)
}

func TestCreateTask_InvalidType(t *testing.T) {
	r := setupTestRouter(t)
	w := httptest.NewRecorder()
	body := map[string]interface{}{ "type": "invalid", "payload": map[string]interface{}{"foo": "bar"} }
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestGetTask_NotFound(t *testing.T) {
	r := setupTestRouter(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/notexist", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}
