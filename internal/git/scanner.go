package git

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// ChangeType represents the kind of Git change observed.
type ChangeType string

const (
	ChangeAdded    ChangeType = "Added"
	ChangeModified ChangeType = "Modified"
	ChangeDeleted  ChangeType = "Deleted"
	ChangeRenamed  ChangeType = "Renamed"
)

// Change represents a single changed file.
type Change struct {
	FileName  string
	Path      string
	Type      ChangeType
	Extension string
}

// Summary contains the counts for each change type.
type Summary struct {
	Added    int
	Modified int
	Deleted  int
	Renamed  int
}

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

	if _, err := git.PlainOpen(absPath); err != nil {
		return "", fmt.Errorf("repository not found at %s", absPath)
	}

	return absPath, nil
}

// Scan compares the current working tree against the provided base branch.
func (s Scanner) Scan(baseBranch string) ([]Change, Summary, error) {
	repoPath, err := s.DetectRepository()
	if err != nil {
		return nil, Summary{}, err
	}

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, Summary{}, fmt.Errorf("unable to open repository: %w", err)
	}

	referenceName := plumbing.NewBranchReferenceName(baseBranch)
	ref, err := repo.Reference(referenceName, true)
	if err != nil {
		return nil, Summary{}, fmt.Errorf("branch %s not found: %w", baseBranch, err)
	}

	if _, err := repo.CommitObject(ref.Hash()); err != nil {
		return nil, Summary{}, fmt.Errorf("unable to resolve commit for branch %s: %w", baseBranch, err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, Summary{}, fmt.Errorf("unable to inspect working tree: %w", err)
	}

	status, err := worktree.Status()
	if err != nil {
		return nil, Summary{}, fmt.Errorf("unable to inspect working tree: %w", err)
	}

	changes := make([]Change, 0, len(status))
	for file, fileStatus := range status {
		changeType := classify(*fileStatus)
		changes = append(changes, Change{
			FileName:  filepath.Base(file),
			Path:      filepath.ToSlash(file),
			Type:      changeType,
			Extension: extension(file),
		})
	}

	summary := Summary{}
	for _, change := range changes {
		switch change.Type {
		case ChangeAdded:
			summary.Added++
		case ChangeModified:
			summary.Modified++
		case ChangeDeleted:
			summary.Deleted++
		case ChangeRenamed:
			summary.Renamed++
		}
	}

	return changes, summary, nil
}

func classify(status git.FileStatus) ChangeType {
	switch status.Staging {
	case git.Added:
		return ChangeAdded
	case git.Modified:
		return ChangeModified
	case git.Deleted:
		return ChangeDeleted
	case git.Renamed:
		return ChangeRenamed
	default:
		return ChangeModified
	}
}

func extension(path string) string {
	base := filepath.Base(path)
	parts := strings.Split(base, ".")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-1]
}
