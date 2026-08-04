package excel

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/OxalcVasquez/git-doc-automation/internal/report"
)

func TestWrite(t *testing.T) {
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "out.xlsx")

	err := Write(outputPath, []report.Entry{{FileName: "main.go", Path: "cmd/main.go", ChangeType: "Modified", Extension: "go", Module: "cmd"}}, report.Summary{Modified: 1})
	if err != nil {
		t.Fatalf("expected write to succeed, got %v", err)
	}

	if _, err := os.Stat(outputPath); err != nil {
		t.Fatalf("expected workbook file to exist, got %v", err)
	}
}
