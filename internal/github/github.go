package github

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v61/github"
)

type Service struct {
	Client *github.Client
}

func NewGitHubService() *Service {
	httpClient := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}

	return &Service{
		Client: github.NewClient(httpClient).WithAuthToken(os.Getenv("GITHUB_TOKEN")),
	}
}

func (s *Service) DownloadRepository(ctx context.Context, repoName string) (*github.Repository, error) {
	parts := strings.Split(repoName, "/")
	if len(parts) != 2 {
		return nil, errors.New("invalid repository name")
	}

	owner, repo := parts[0], parts[1]
	
	archiveURL, _, err := s.Client.Repositories.GetArchiveLink(ctx, owner, repo, github.Tarball, nil, 3)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, archiveURL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	
}