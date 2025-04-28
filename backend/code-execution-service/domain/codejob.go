// domain/codejob.go
package domain

type CodeJob struct {
	ID       string `json:"id"`
	TaskID   string `json:"task_id"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
	Output   string `json:"output"`
	ExitCode int    `json:"exit_code"`
	Status   string `json:"status"`
}
