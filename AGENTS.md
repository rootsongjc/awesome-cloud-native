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

- **`scripts/check_status.py`**: Parses README for GitHub URLs → queries GitHub API → writes `status.json`. Needs `GITHUB_TOKEN` (5000 req/hour vs 60 unauthenticated). Inactive threshold: 2 years since last push. Retries on API rate limits by sleeping until the limit reset, so throttled repos are never silently dropped.
- **`tmpl/assets/status-checker.js`**: Frontend JS fetches `status.json`, injects colored dots before project links (green=active, yellow=archived, red=deleted, gray=inactive) and a summary banner at top.
- **`.github/workflows/check-status.yml`**: Runs on README pushes, weekly Monday 00:03 UTC, and manual trigger. Runs are serialized through a `check-status` concurrency group (newer runs cancel older ones). Copies `status.json` to `tmpl/` and opens a bot PR titled "Update project status".

Status flow: `README.md → Python script → status.json (bot PR) → frontend JS renders dots`

### CI Workflows

- **`.github/workflows/verify-site.yml`**: On PRs, runs `go test ./...` (ordering and duplicate checks) and rejects hand edits to generated files (`tmpl/index.html`, `tmpl/status.json`) from non-bot authors. On pushes to main, regenerates `tmpl/index.html` via `make build` and pushes the result as `github-actions[bot]`.
- Contributors only edit `README.md`. The generated site and status data are maintained entirely by CI.

## Conventions

- **Alphabetical order** within each category. `go test ./...` enforces this.
- **One entry per project**, single best category.
- **Inclusion criteria**: OSI-approved license, canonical upstream repo, cloud-native relevance. Full rules in `CONTRIBUTING.md` and the issue/PR templates under `.github/`.
- **Descriptions**: single line, sentence case, no trailing periods unless complete sentence.
- **Link text**: official project name in Title Case, pointing to canonical GitHub repo.
- **Commit messages**: imperative tone (`Add litellm`, `Bump golang.org/x/net`).
- **PR workflow**: edit `README.md` only → `go test ./...` → open PR. CI regenerates `tmpl/index.html` after merge; never commit generated files.
