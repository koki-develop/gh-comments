package main

import (
	"os"

	"github.com/koki-develop/gh-comments/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
