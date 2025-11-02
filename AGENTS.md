# Repository Guidelines

## Project Structure & Module Organization

The curated catalog lives in `README.md`, arranged by single-level categories. Supporting Go code (`repo.go`, `repo_test.go`) transforms that Markdown into the static site under `tmpl/`. Generated assets are written to `tmpl/index.html` using `tmpl/tmpl.html` as the layout and `tmpl/assets/` for shared resources. `CONTRIBUTING.md` describes community expectations; treat it as a companion reference when updating the list.

## Build, Test, and Development Commands

- `go run ./repo.go`: Regenerates `tmpl/index.html` from the current README content. Use this before publishing to confirm the site view renders correctly.
- `go test ./...`: Runs the verification suite that enforces alphabetical grouping and guards against duplicate links in the README.
- `gofmt -w repo.go repo_test.go`: Normalizes Go source formatting whenever edits are made.

## Coding Style & Naming Conventions

Maintain alphabetical ordering within each Markdown category and keep link text in Title Case (mirroring upstream project names). Descriptions should stay on a single line, use sentence case, and avoid trailing periods unless the sentence is complete. Go code follows standard `gofmt` conventions, tabs for indentation, and short, descriptive function names (`updateRepo`, `readMarkdownFile`). Template IDs and filenames remain lowercase-with-dashes.

## Testing Guidelines

`go test ./...` must pass before submission; it parses the README to ensure every list item remains alphabetized and unique. When adding a project, re-run the tests after inserting the entry and adjust placement until the suite succeeds. If you modify the generator or templates, add focused unit tests in `repo_test.go` to cover the new behavior and keep fixtures minimal.

## Commit & Pull Request Guidelines

Follow the existing Git history’s imperative tone (e.g., `Add litellm`, `Bump golang.org/x/net`). Each commit should group related changes—avoid mixing list edits with generator updates. Pull requests need a concise summary, the commands executed (`go test ./...`), and any context about category decisions. Include screenshots only when UI output in `tmpl/index.html` materially changes. Reference related issues or upstream announcements where applicable to justify additions or removals.

## Curation Workflow Tips

Before adding an item, verify the repository is active and open source, and confirm it fits exactly one category. Prefer placing new tools under the most specific heading introduced in the reorganized taxonomy (e.g., `Build & Packaging Automation` rather than the broader `Continuous Delivery & GitOps`). When in doubt, document the rationale in the PR to help maintainers keep the list consistent over time.
