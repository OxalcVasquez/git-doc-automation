# git-doc-automation

A Go-based educational project that automates release documentation from Git changes.

## What it does

The tool scans a Git repository, groups changed files by type, and generates an Excel workbook with:

- a detail sheet for the changed files
- a summary sheet with counts for Added, Modified, Deleted, and Renamed files

## Project structure

- cmd/: application entrypoints
- internal/cli/: command-line configuration and orchestration
- internal/git/: Git repository scanning and change detection
- internal/report/: report modeling and summary generation
- internal/excel/: Excel workbook export
- scripts/: helper script entrypoints
- output/: generated Excel files
- docs/: project documentation

## Run it

```bash
go run ./cmd --base-branch main
```

Optional flags:

```bash
go run ./cmd --base-branch main --output ./output/documentation.xlsx --repo .
```

## Output

The tool writes the Excel file to the configured output path, by default:

- ./output/documentation.xlsx
