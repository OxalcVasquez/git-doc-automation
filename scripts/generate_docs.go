package main

import (
	"fmt"
	"os"

	"github.com/OxalcVasquez/git-doc-automation/internal/cli"
)

func main() {
	cfg, err := cli.ParseConfig(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := cli.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
