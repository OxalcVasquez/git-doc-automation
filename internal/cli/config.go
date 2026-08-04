package cli

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

// Config holds the runtime options for the documentation generator.
type Config struct {
	BaseBranch string
	OutputPath string
	RepoPath   string
}

// ParseConfig parses CLI flags into a Config value.
func ParseConfig(args []string) (Config, error) {
	fs := flag.NewFlagSet("git-doc-automation", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	cfg := Config{
		BaseBranch: "main",
		OutputPath: "./output/documentation.xlsx",
		RepoPath:   ".",
	}

	fs.StringVar(&cfg.BaseBranch, "base-branch", cfg.BaseBranch, "branch to compare against")
	fs.StringVar(&cfg.OutputPath, "output", cfg.OutputPath, "path to write the generated Excel file")
	fs.StringVar(&cfg.RepoPath, "repo", cfg.RepoPath, "path to the git repository")

	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}

	cfg.BaseBranch = strings.TrimSpace(cfg.BaseBranch)
	cfg.OutputPath = strings.TrimSpace(cfg.OutputPath)
	cfg.RepoPath = strings.TrimSpace(cfg.RepoPath)

	if cfg.BaseBranch == "" {
		return Config{}, fmt.Errorf("base branch cannot be empty")
	}

	return cfg, nil
}
