package cmd

import (
	"github.com/spf13/cobra"
)

var repoFlag string

var rootCmd = &cobra.Command{
	Use:   "gh-comments",
	Short: "GitHub CLI extension for managing issue and PR comments",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&repoFlag, "repo", "R", "", "Select another repository using the [HOST/]OWNER/REPO format")
}
