package git

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectRsepository(t *testing.T) {
	tempDir := t.TempDir()
	gitDir := filepath.Join(tempDir, ".git")
	if err := os.Mkdir(gitDir, 0o755); err != nil {
		t.Fatalf("failed to create .git dir: %v", err)
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
