// domain/task.go
package domain

type Task struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	Payload   string `json:"payload"`
	Result    string `json:"result"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Step struct {
	ID       string `json:"id"`
	TaskID   string `json:"task_id"`
	Type     string `json:"type"`
	Input    string `json:"input"`
	Output   string `json:"output"`
	Sequence int    `json:"sequence"`
	Status   string `json:"status"`
}
