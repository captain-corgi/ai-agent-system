// domain/plan.go
package domain

type Plan struct {
	ID      string   `json:"id"`
	TaskID  string   `json:"task_id"`
	Steps   []Step   `json:"steps"`
}

type Step struct {
	ID       string `json:"id"`
	PlanID   string `json:"plan_id"`
	Type     string `json:"type"`
	Input    string `json:"input"`
	Output   string `json:"output"`
	Sequence int    `json:"sequence"`
	Status   string `json:"status"`
}
