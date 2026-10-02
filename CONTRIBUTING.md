# Contribution Guidelines

Thanks for helping keep this Kubernetes and cloud-native catalog sharp. Please follow the rules below so the list stays consistent and searchable.

## Inclusion criteria

Before suggesting a project, check it against these rules. Submissions that fail any of them will be declined:

- **Open source license.** The repository must carry an OSI-approved license (Apache-2.0, MIT, BSD, MPL, GPL family, etc.). Source-available licenses such as PolyForm, BSL, or SSPL do **not** qualify, even for "community editions".
- **Canonical repository.** Link to the upstream project itself, not a fork or a mirror. If the upstream is already listed, don't add forks.
- **Cloud-native relevance.** The project must clearly benefit cloud-native or Kubernetes workloads. Quick test: would a platform/DevOps engineer reach for it when building or operating cloud-native systems?
- **Maintained or established.** Prefer actively maintained projects. The status checker flags repos with no push in 2 years as inactive; historically significant but dormant projects may stay, but brand-new unmaintained ones won't be added.
- **Vendor-affiliated is fine.** A tool maintained by a company qualifies as long as the listed artifact itself is genuinely open source and standalone.

Not sure? Open an issue using the *Suggest a project* template and ask. That's what it's for.

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

You only need to edit `README.md`. Do not commit `tmpl/index.html` or `tmpl/status.json` – they are generated files, rebuilt automatically by CI after merge.

Include the commands you ran in your pull request description.

## Reporting issues

Open an issue when you find outdated, broken, or abandoned projects, or when you’d like to propose a new category. Provide links or context so maintainers can validate quickly.

Thanks again for contributing!
