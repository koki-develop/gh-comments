package api

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/koki-develop/gh-comments/internal/models"
)

type Client struct {
	rest *api.RESTClient
}

func NewClient() (*Client, error) {
	rest, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}
	return &Client{rest: rest}, nil
}

func (c *Client) ListComments(owner, repo string, issueNumber int) ([]models.Comment, error) {
	path := fmt.Sprintf("repos/%s/%s/issues/%d/comments", owner, repo, issueNumber)
	var comments []models.Comment
	if err := c.rest.Get(path, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}
