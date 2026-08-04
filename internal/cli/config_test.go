package cli
package cli

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	cfg, err := ParseConfig([]string{"--base-branch", "develop", "--output", "./tmp/report.xlsx"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	want := Config{
		BaseBranch: "develop",
		OutputPath: "./tmp/report.xlsx",
		RepoPath:   ".",
	}

	if !reflect.DeepEqual(cfg, want) {
		t.Fatalf("unexpected config: got %+v want %+v", cfg, want)
	}
}
