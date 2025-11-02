# Contribution Guidelines

Thanks for helping keep this Kubernetes and cloud-native catalog sharp. Please follow the rules below so the list stays consistent and searchable.

## General expectations

- Additions **must be open source** (ideally GitHub-hosted) and clearly benefit cloud-native workloads.
- Place each project in the **single best matching top-level category** from the table of contents. If unsure, explain the rationale in the PR description.
- Keep descriptions on a single line, start with sentence case, and avoid marketing fluff.
- Use the official project name as the link text and point directly to the canonical repository.

## Ordering & formatting

- Categories and items are maintained in **alphabetical order**. Run `go test ./...` (see below) before submitting; it will fail if ordering or duplicates are wrong.
- One entry per project. Remove retired or duplicate links when encountered.
- New entries go directly under the relevant `##` heading with a trailing period only if the sentence is complete.

## Local checks before PR

1. `gofmt -w repo.go repo_test.go` (only needed if you touched Go files).
2. `go test ./...` – verifies alphabetical ordering and ensures no duplicate links.
3. `go run ./repo.go` – regenerates `tmpl/index.html`; run this whenever README content changes so the published site stays in sync.

Include the commands you ran in your pull request description.

## Reporting issues

Open an issue when you find outdated, broken, or abandoned projects, or when you’d like to propose a new category. Provide links or context so maintainers can validate quickly.

Thanks again for contributing!
