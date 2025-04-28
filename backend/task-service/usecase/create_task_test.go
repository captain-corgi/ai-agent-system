package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/captain-corgi/ai-agent-system/task-service/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name     string
		typeArg  string
		payload  interface{}
		repoErr  error
		expectErr bool
	}{
		{"valid_code", "code", "print('hi')", nil, false},
		{"valid_browse", "browse", "url", nil, false},
		{"invalid_type", "invalid", "foo", nil, true},
		{"repo_error", "code", "fail", errors.New("db fail"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mockTaskRepo)
			if tt.name != "invalid_type" {
				repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(tt.repoErr)
			}
			task, err := CreateTask(context.Background(), repo, tt.typeArg, tt.payload)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, task)
			}
		})
	}
}

func TestListTasks(t *testing.T) {
	repo := new(mockTaskRepo)
	repo.On("ListAll", mock.Anything).Return([]*domain.Task{{ID: "t1"}}, nil)
	tasks, err := ListTasks(context.Background(), repo)
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
}

func TestGetTask(t *testing.T) {
	repo := new(mockTaskRepo)
	repo.On("GetByID", mock.Anything, "tid").Return(&domain.Task{ID: "tid"}, nil)
	task, err := GetTask(context.Background(), repo, "tid")
	assert.NoError(t, err)
	assert.Equal(t, "tid", task.ID)
}
