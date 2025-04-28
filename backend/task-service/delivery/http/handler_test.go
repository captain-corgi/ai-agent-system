package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/captain-corgi/ai-agent-system/task-service/domain"
)

type mockTaskRepo struct {
	mock.Mock
}

func (m *mockTaskRepo) Create(ctx context.Context, task *domain.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}
func (m *mockTaskRepo) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Task), args.Error(1)
}
func (m *mockTaskRepo) ListAll(ctx context.Context) ([]*domain.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func TestCreateTaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := new(mockTaskRepo)
	h := NewHandler(repo)
	r := gin.Default()
	h.RegisterRoutes(r)

	// Valid request
	reqBody := map[string]interface{}{"type": "code", "payload": "print('hi')"}
	jsonBody, _ := json.Marshal(reqBody)
	repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(nil).Once()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Invalid type
	reqBody = map[string]interface{}{"type": "invalid", "payload": "foo"}
	jsonBody, _ = json.Marshal(reqBody)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := new(mockTaskRepo)
	h := NewHandler(repo)
	r := gin.Default()
	h.RegisterRoutes(r)

	task := &domain.Task{ID: "123", Type: "code"}
	repo.On("GetByID", mock.Anything, "123").Return(task, nil).Once()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/123", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Not found
	repo.On("GetByID", mock.Anything, "notfound").Return((*domain.Task)(nil), nil).Once()
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/tasks/notfound", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListTasksHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := new(mockTaskRepo)
	h := NewHandler(repo)
	r := gin.Default()
	h.RegisterRoutes(r)

	tasks := []*domain.Task{{ID: "1"}, {ID: "2"}}
	repo.On("ListAll", mock.Anything).Return(tasks, nil).Once()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
