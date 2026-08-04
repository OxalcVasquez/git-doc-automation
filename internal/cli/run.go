package cli

import (
	"fmt"

	"github.com/OxalcVasquez/git-doc-automation/internal/git"
	"github.com/OxalcVasquez/git-doc-automation/internal/report"
)

// Run executes the current CLI flow with the given configuration.
func Run(cfg Config) error {
	fmt.Printf("Configuration loaded\n")
	fmt.Printf("Base branch: %s\n", cfg.BaseBranch)
	fmt.Printf("Output path: %s\n", cfg.OutputPath)
	fmt.Printf("Repository path: %s\n", cfg.RepoPath)

	scanner := git.NewScanner(cfg.RepoPath)
	if _, err := scanner.DetectRepository(); err != nil {
		return err
	}

	changes, summary, err := scanner.Scan(cfg.BaseBranch)
	if err != nil {
		return err
	}

	entries, reportSummary := report.Build(changes, summary)

	fmt.Println("------------------------------------")
	fmt.Println("Repository scanned")
	fmt.Printf("Added: %d\n", reportSummary.Added)
	fmt.Printf("Modified: %d\n", reportSummary.Modified)
	fmt.Printf("Deleted: %d\n", reportSummary.Deleted)
	fmt.Printf("Renamed: %d\n", reportSummary.Renamed)
	fmt.Println("------------------------------------")
	fmt.Printf("Total changes: %d\n", len(entries))

	return nil
}
