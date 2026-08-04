package excel

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/OxalcVasquez/git-doc-automation/internal/report"
	"github.com/xuri/excelize/v2"
)

// Write creates an Excel workbook with a details sheet and a summary sheet.
func Write(outputPath string, entries []report.Entry, summary report.Summary) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	f := excelize.NewFile()
	defer func() {
		_ = f.Close()
	}()

	if err := f.SetSheetRow("Sheet1", "A1", &[]string{"File Name", "Path", "Change Type", "Extension", "Module"}); err != nil {
		return fmt.Errorf("write headers: %w", err)
	}

	for i, entry := range entries {
		row := i + 2
		values := []string{entry.FileName, entry.Path, entry.ChangeType, entry.Extension, entry.Module}
		if err := f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", row), &values); err != nil {
			return fmt.Errorf("write row %d: %w", row, err)
		}
	}

	if _, err := f.NewSheet("Summary"); err != nil {
		return fmt.Errorf("create summary sheet: %w", err)
	}

	if err := f.SetSheetRow("Summary", "A1", &[]string{"Metric", "Value"}); err != nil {
		return fmt.Errorf("write summary headers: %w", err)
	}

	rows := []struct {
		label string
		value string
	}{
		{label: "Added", value: fmt.Sprintf("%d", summary.Added)},
		{label: "Modified", value: fmt.Sprintf("%d", summary.Modified)},
		{label: "Deleted", value: fmt.Sprintf("%d", summary.Deleted)},
		{label: "Renamed", value: fmt.Sprintf("%d", summary.Renamed)},
	}

	for i, row := range rows {
		values := []string{row.label, row.value}
		if err := f.SetSheetRow("Summary", fmt.Sprintf("A%d", i+2), &values); err != nil {
			return fmt.Errorf("write summary row %d: %w", i+2, err)
		}
	}

	f.SetActiveSheet(2)

	if err := f.SaveAs(outputPath); err != nil {
		return fmt.Errorf("save workbook: %w", err)
	}

	return nil
}
