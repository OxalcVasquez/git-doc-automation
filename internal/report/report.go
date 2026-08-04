package report

import "github.com/OxalcVasquez/git-doc-automation/internal/git"

// Entry represents a single file entry to be rendered in the report.
type Entry struct {
	FileName   string
	Path       string
	ChangeType string
	Extension  string
	Module     string
}

// Summary contains the high-level counts for the report.
type Summary struct {
	Added    int
	Modified int
	Deleted  int
	Renamed  int
}

// Build creates a report entry list and summary from scanned Git changes.
func Build(changes []git.Change, summary git.Summary) ([]Entry, Summary) {
	entries := make([]Entry, 0, len(changes))
	for _, change := range changes {
		entries = append(entries, Entry{
			FileName:   change.FileName,
			Path:       change.Path,
			ChangeType: string(change.Type),
			Extension:  change.Extension,
			Module:     moduleName(change.Path),
		})
	}

	return entries, Summary{
		Added:    summary.Added,
		Modified: summary.Modified,
		Deleted:  summary.Deleted,
		Renamed:  summary.Renamed,
	}
}

func moduleName(path string) string {
	if path == "" {
		return "root"
	}

	parts := []rune(path)
	for i, part := range parts {
		if part == '/' {
			return string(parts[:i])
		}
	}

	return "root"
}
