# Contribution Guidelines

Thanks for helping keep this Kubernetes and cloud-native catalog sharp. Please follow the rules below so the list stays consistent and searchable.

## Inclusion criteria

Not every submission is accepted. A project has to clear all three groups below.

### Hard requirements

Failing any of these means the project cannot be listed:

- **Open source license.** The repository must carry an OSI-approved license (Apache-2.0, MIT, BSD, MPL, GPL family, etc.). Source-available licenses such as PolyForm, BSL, or SSPL do **not** qualify, even for "community editions".
- **Canonical repository.** Link to the public upstream project itself, not a fork, mirror, or a URL that redirects elsewhere.
- **Not already listed.** One entry per project, in a single category.
- **Not archived.** Archived or deleted repositories are not added.

### Maturity thresholds

The list is a catalog of projects people can depend on, not a launch channel. A project that meets the hard requirements but not these is marked *too early* and is welcome back later:

- **Adoption:** at least **50 GitHub stars**.
- **Age:** the repository is at least **90 days** old.
- **Release:** at least one tagged release (not required under *Tutorials & Learning*).
- **Activity:** a push within the last **12 months**.

Maintainers may waive a threshold for projects with clear standing, such as CNCF or other foundation-hosted projects.

### Relevance and quality

These are judgement calls made by a maintainer:

- **Cloud-native relevance.** The project must clearly benefit cloud-native or Kubernetes workloads. Quick test: would a platform/DevOps engineer reach for it when building or operating cloud-native systems?
- **Standalone and usable.** It has documentation that lets someone else install and use it. Vendor-affiliated tools are fine as long as the listed artifact is genuinely open source and useful without the vendor's paid product.
- **Best category.** It fits one existing top-level category.

### How submissions are reviewed

Every *Suggest a project* issue and every pull request that adds an entry is checked automatically by [@jimmysongio](https://github.com/jimmysongio), the maintainer's AI agent. It posts one comment with the result of each check and sets a label:

| Label | Meaning |
| --- | --- |
| `review/eligible` | Passes the automated checks; waiting for a maintainer's decision. |
| `review/needs-changes` | Fixable problems (link, name, description). Fix them and comment `/review`. |
| `review/too-early` | Below the maturity thresholds. Resubmit when the project meets them. |
| `review/not-eligible` | Fails a hard requirement. |
| `review/needs-maintainer` | Could not be decided automatically, e.g. an unrecognised license. |

The agent never merges or closes anything. A maintainer makes the final call and may override any result. Comment `/review` on your issue or pull request to re-run the check.

Not sure? Open an issue using the *Suggest a project* template and ask. That's what it's for.

## General expectations

- Additions **must be open source** (ideally GitHub-hosted) and clearly benefit cloud-native workloads.
- Place each project in the **single best matching top-level category** from the table of contents. If unsure, explain the rationale in the PR description.
- Write descriptions in English, even if the upstream project documents itself in another language. Keep them on a single line, start with sentence case, and avoid marketing fluff.
- Use the project name in lowercase, without spaces or slashes, as the link text (e.g. `bank-vaults`, not `banzaicloud/bank-vaults`) and point directly to the canonical repository.

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
