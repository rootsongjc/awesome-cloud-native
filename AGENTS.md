# Repository Guidelines

## Build & Dev Commands

```bash
make build     # Generate tmpl/index.html from README.md (go run repo.go)
make preview   # Build + start local HTTP server on tmpl/ (default port 8080, override: PORT=3000)
make test      # Run Go tests (alphabetical order + duplicate link checks)
make status    # Run project status checker, outputs status.json and copies to tmpl/
make clean     # Remove generated files
```

Formatting: `gofmt -w repo.go repo_test.go`

## Architecture

This is a curated "Awesome" list of cloud native projects. The static site is generated from a single source of truth:

```plaintext
README.md → repo.go (Go + GFM) → tmpl/index.html → served as static site
```

- **`README.md`**: The catalog. ~800 projects in ~34 single-level categories. Format: `- [name](https://github.com/owner/repo) - Description.`
- **`repo.go`**: Reads README.md, converts with `github_flavored_markdown`, renders via `tmpl/tmpl.html`, writes `tmpl/index.html`. Also auto-pulls git changes.
- **`repo_test.go`**: Parses rendered HTML to enforce alphabetical ordering within each category and reject duplicate links.
- **`tmpl/`**: Static site output. CSS/JS assets in `tmpl/assets/`. Do not manually edit `index.html`.

### Project Status Tracking

Automated health check for all listed GitHub repos:

- **`scripts/check_status.py`**: Parses README for GitHub URLs → queries GitHub API → writes `status.json`. Needs `GITHUB_TOKEN` (5000 req/hour vs 60 unauthenticated). Inactive threshold: 2 years since last push.
- **`tmpl/assets/status-checker.js`**: Frontend JS fetches `status.json`, injects colored dots before project links (green=active, yellow=archived, red=deleted, gray=inactive) and a summary banner at top.
- **`.github/workflows/check-status.yml`**: Runs weekly Monday 00:03 UTC + manual trigger. Creates PR with updated `status.json` (both repo root and `tmpl/`). Does NOT auto-push to main.

Status flow: `README.md → Python script → status.json (PR) → frontend JS renders dots`

## Conventions

- **Alphabetical order** within each category. `go test ./...` enforces this.
- **One entry per project**, single best category.
- **Descriptions**: single line, sentence case, no trailing periods unless complete sentence.
- **Link text**: official project name in Title Case, pointing to canonical GitHub repo.
- **Commit messages**: imperative tone (`Add litellm`, `Bump golang.org/x/net`).
- **PR workflow**: add entry → `go test ./...` → `make build` → include commands in PR description.
