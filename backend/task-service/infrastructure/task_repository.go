// infrastructure/task_repository.go
package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/captain-corgi/ai-agent-system/task-service/domain"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	GetByID(ctx context.Context, id string) (*domain.Task, error)
	ListAll(ctx context.Context) ([]*domain.Task, error)
}

type SQLiteTaskRepository struct {
	db *sql.DB
}

func NewSQLiteTaskRepository(db *sql.DB) TaskRepository {
	return &SQLiteTaskRepository{db: db}
}

func (r *SQLiteTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	query := `INSERT INTO tasks (id, type, status, payload, result, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, task.ID, task.Type, task.Status, task.Payload, task.Result, task.CreatedAt, task.UpdatedAt)
	return err
}

func (r *SQLiteTaskRepository) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	query := `SELECT id, type, status, payload, result, created_at, updated_at FROM tasks WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)
	task := &domain.Task{}
	err := row.Scan(&task.ID, &task.Type, &task.Status, &task.Payload, &task.Result, &task.CreatedAt, &task.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *SQLiteTaskRepository) ListAll(ctx context.Context) ([]*domain.Task, error) {
	query := `SELECT id, type, status, payload, result, created_at, updated_at FROM tasks`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*domain.Task{}
	for rows.Next() {
		task := &domain.Task{}
		if err := rows.Scan(&task.ID, &task.Type, &task.Status, &task.Payload, &task.Result, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Utility for DB migration (table creation)
func Migrate(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		type TEXT,
		status TEXT,
		payload TEXT,
		result TEXT,
		created_at TEXT,
		updated_at TEXT
	)`
	_, err := db.Exec(query)
	return err
}
