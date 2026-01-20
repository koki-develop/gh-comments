# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

gh-comments is a GitHub CLI extension for managing issue and PR comments. It uses the `gh` CLI's extension system and the `go-gh` library for GitHub API interactions.

## Build & Development Commands

```bash
# Build the extension
go build -o gh-comments .

# Run locally (after build)
./gh-comments list <issue-number>

# Install as gh extension (for testing)
gh extension install .

# Run as gh extension
gh comments list <issue-number>
```

## Architecture

- **main.go** - Entry point
- **cmd/** - Cobra command definitions
  - `root.go` - Root command with global flags (e.g., `--repo`)
  - `list.go` - List comments command implementation
- **internal/api/** - GitHub API client wrapper using `go-gh`
- **internal/models/** - Data structures (Comment, User)

The extension uses `github.com/cli/go-gh/v2` for GitHub API access and authentication, inheriting credentials from the `gh` CLI.

## Release

Releases are automated via GitHub Actions using `cli/gh-extension-precompile` when a version tag (v*) is pushed.
