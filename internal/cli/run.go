package cli

import "fmt"

// Run executes the current CLI flow with the given configuration.
func Run(cfg Config) error {
	fmt.Printf("Configuration loaded\n")
	fmt.Printf("Base branch: %s\n", cfg.BaseBranch)
	fmt.Printf("Output path: %s\n", cfg.OutputPath)
	fmt.Printf("Repository path: %s\n", cfg.RepoPath)
	return nil
}
