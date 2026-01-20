package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/koki-develop/gh-comments/internal/api"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list <issue-number>",
	Short: "List comments on an issue or pull request",
	Args:  cobra.ExactArgs(1),
	RunE:  runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) error {
	issueNumber, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid issue number: %s", args[0])
	}

	owner, repo, err := resolveRepository()
	if err != nil {
		return err
	}

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	comments, err := client.ListComments(owner, repo, issueNumber)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(comments)
}

func resolveRepository() (owner, repo string, err error) {
	if repoFlag != "" {
		parts := strings.Split(repoFlag, "/")
		if len(parts) == 2 {
			return parts[0], parts[1], nil
		}
		if len(parts) == 3 {
			return parts[1], parts[2], nil
		}
		return "", "", fmt.Errorf("invalid repository format: %s (expected [HOST/]OWNER/REPO)", repoFlag)
	}

	r, err := repository.Current()
	if err != nil {
		return "", "", fmt.Errorf("could not determine repository: %w\nUse --repo flag to specify the repository", err)
	}

	return r.Owner, r.Name, nil
}
