package report

import (
	"testing"

	"github.com/OxalcVasquez/git-doc-automation/internal/git"
)

func TestBuild(t *testing.T) {
	changes := []git.Change{{
		FileName:  "main.go",
		Path:      "cmd/main.go",
		Type:      git.ChangeModified,
		Extension: "go",
	}}

	entries, summary := Build(changes, git.Summary{Modified: 1})

	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}

	if entries[0].Module != "cmd" {
		t.Fatalf("expected module cmd, got %q", entries[0].Module)
	}

	if summary.Modified != 1 {
		t.Fatalf("expected modified count 1, got %d", summary.Modified)
	}
}
