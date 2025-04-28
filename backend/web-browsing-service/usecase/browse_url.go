package usecase

import (
	"errors"
	"fmt"
	"github.com/captain-corgi/ai-agent-system/web-browsing-service/domain"
)

// BrowseURL mocks browsing a URL. Replace with chromedp integration for real browsing.
func BrowseURL(url string) (*domain.BrowseJob, error) {
	if url == "" {
		return nil, errors.New("url is required")
	}
	job := &domain.BrowseJob{
		ID:     "browsejob-1",
		URL:    url,
		Status: "completed",
		Result: fmt.Sprintf("Browsed %s", url),
	}
	return job, nil
}
