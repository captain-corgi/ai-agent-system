// domain/fsjob.go
package domain

type FsJob struct {
	ID      string `json:"id"`
	TaskID  string `json:"task_id"`
	OpType  string `json:"op_type"`
	Path    string `json:"path"`
	Content string `json:"content"`
	Result  string `json:"result"`
	Status  string `json:"status"`
}
