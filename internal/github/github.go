package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v61/github"
)

type GitHubClient struct {
	*github.Client
}

func NewGitHubClient(token string) *GitHubClient {
	client := github.NewClient(nil).WithAuthToken(token)
	return &GitHubClient{
		Client: client,
	}
}

func (c *GitHubClient) ListOrgRepos(ctx context.Context, org string) ([]*github.Repository, error) {
	var repoList []*github.Repository

	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	repoChan := make(chan []*github.Repository)
	errChan := make(chan error)

	go func() {
		defer close(repoChan)

		for {
			repos, resp, err := c.Repositories.ListByOrg(ctx, org, opts)
			if err != nil {
				errChan <- fmt.Errorf("failed to list repositories: %w", err)
				return
			}

			repoChan <- repos

			if resp.NextPage == 0 {
				break
			}
			opts.Page = resp.NextPage
		}
	}()

	for repos := range repoChan {
		repoList = append(repoList, repos...)
	}

	select {
	case err := <-errChan:
		return nil, err
	default:
		return repoList, nil
	}
}
