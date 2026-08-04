package cli

import (
	"fmt"

	"github.com/OxalcVasquez/git-doc-automation/internal/git"
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

	fmt.Println("------------------------------------")
	fmt.Println("Repository scanned")
	fmt.Printf("Added: %d\n", summary.Added)
	fmt.Printf("Modified: %d\n", summary.Modified)
	fmt.Printf("Deleted: %d\n", summary.Deleted)
	fmt.Printf("Renamed: %d\n", summary.Renamed)
	fmt.Println("------------------------------------")
	fmt.Printf("Total changes: %d\n", len(changes))

	return nil
}
