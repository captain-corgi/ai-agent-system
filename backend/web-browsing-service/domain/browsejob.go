// domain/browsejob.go
package domain

type BrowseJob struct {
	ID      string `json:"id"`
	TaskID  string `json:"task_id"`
	URL     string `json:"url"`
	Result  string `json:"result"`
	Status  string `json:"status"`
}
