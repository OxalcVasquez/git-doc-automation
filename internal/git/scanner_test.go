package git

import (
	"testing"

	"github.com/go-git/go-git/v5"
)

func TestDetectRsepository(t *testing.T) {
	tempDir := t.TempDir()
	if _, err := git.PlainInit(tempDir, false); err != nil {
		t.Fatalf("failed to initialize git repo: %v", err)
	}

	scanner := NewScanner(tempDir)
	got, err := scanner.DetectRepository()
	if err != nil {
		t.Fatalf("expected repository detection to succeed, got %v", err)
	}

	if got != tempDir {
		t.Fatalf("expected repository path %q, got %q", tempDir, got)
	}
}

func TestClassify(t *testing.T) {
	tests := []struct {
		name   string
		status git.FileStatus
		want   ChangeType
	}{
		{name: "added", status: git.FileStatus{Staging: git.Added}, want: ChangeAdded},
		{name: "modified", status: git.FileStatus{Staging: git.Modified}, want: ChangeModified},
		{name: "deleted", status: git.FileStatus{Staging: git.Deleted}, want: ChangeDeleted},
		{name: "renamed", status: git.FileStatus{Staging: git.Renamed}, want: ChangeRenamed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := classify(tt.status); got != tt.want {
				t.Fatalf("classify(%v) = %q, want %q", tt.status.Staging, got, tt.want)
			}
		})
	}
}
