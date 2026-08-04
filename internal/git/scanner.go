package git

import (
	"fmt"
	"os"
	"path/filepath"
)

// Scanner inspects a repository and collects change information.
type Scanner struct {
	RepoPath string
}

// NewScanner creates a scanner for the provided repository path.
func NewScanner(repoPath string) Scanner {
	return Scanner{RepoPath: repoPath}
}

// DetectRepository verifies that the target path is a Git repository.
func (s Scanner) DetectRepository() (string, error) {
	absPath, err := filepath.Abs(s.RepoPath)
	if err != nil {
		return "", err
	}

	gitDir := filepath.Join(absPath, ".git")
	if _, err := os.Stat(gitDir); err != nil {
		return "", fmt.Errorf("repository not found at %s", absPath)
	}

	return absPath, nil
}
