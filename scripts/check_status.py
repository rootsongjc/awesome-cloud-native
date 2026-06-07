#!/usr/bin/env python3
"""Check GitHub project status for awesome-cloud-native README.

Parses README.md for GitHub repo URLs, queries the GitHub API for each,
and outputs status.json with project health data.
"""

import json
import os
import re
import sys
import urllib.request
import urllib.error
from datetime import datetime, timezone, timedelta

README_PATH = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), "README.md")
STATUS_PATH = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), "status.json")

# Match: - [name](https://github.com/owner/repo)
GITHUB_LINK_RE = re.compile(r"\[.*?\]\(https://github\.com/([^/)]+/[^/)]+?)(?:[/?].*)?\)")

# Status thresholds
INACTIVE_THRESHOLD_DAYS = 730  # 2 years
API_BATCH_SIZE = 50


def parse_github_repos(readme_path):
    """Extract unique owner/repo strings from README markdown links."""
    with open(readme_path, "r", encoding="utf-8") as f:
        content = f.read()

    repos = set()
    for match in GITHUB_LINK_RE.finditer(content):
        repo = match.group(1)
        if "/" in repo:
            repos.add(repo)

    return sorted(repos)


def query_repo_status(repo, token):
    """Query a single repo's status from GitHub API.

    Returns dict with: status, archived, pushed_at
    """
    url = "https://api.github.com/repos/" + repo
    headers = {
        "Accept": "application/vnd.github.v3+json",
        "User-Agent": "awesome-cloud-native-status-checker",
    }
    if token:
        headers["Authorization"] = "Bearer " + token

    req = urllib.request.Request(url, headers=headers)

    try:
        with urllib.request.urlopen(req, timeout=30) as resp:
            data = json.loads(resp.read().decode("utf-8"))
            pushed_at = data.get("pushed_at")
            archived = data.get("archived", False)

            if archived:
                return {"status": "archived", "archived": True, "pushed_at": pushed_at}

            if pushed_at:
                last_push = datetime.fromisoformat(pushed_at.replace("Z", "+00:00"))
                cutoff = datetime.now(timezone.utc) - timedelta(days=INACTIVE_THRESHOLD_DAYS)
                if last_push < cutoff:
                    return {"status": "inactive", "archived": False, "pushed_at": pushed_at}

            return {"status": "active", "archived": False, "pushed_at": pushed_at}

    except urllib.error.HTTPError as e:
        if e.code == 404:
            return {"status": "deleted", "archived": False, "pushed_at": None}
        print("Warning: HTTP " + str(e.code) + " for " + repo, file=sys.stderr)
        return None
    except Exception as e:
        print("Warning: " + str(e) + " for " + repo, file=sys.stderr)
        return None


def check_all_repos(repos, token):
    """Query all repos and return status dict."""
    projects = {}
    total = len(repos)

    for i, repo in enumerate(repos):
        if i > 0 and i % API_BATCH_SIZE == 0:
            print("Progress: " + str(i) + "/" + str(total) + " repos checked...")

        result = query_repo_status(repo, token)
        if result is not None:
            projects[repo] = result

    return projects


def build_status_json(projects):
    """Build the final status.json structure."""
    summary = {"active": 0, "archived": 0, "deleted": 0, "inactive": 0}

    for info in projects.values():
        status = info["status"]
        if status in summary:
            summary[status] += 1

    return {
        "updated_at": datetime.now(timezone.utc).strftime("%Y-%m-%dT%H:%M:%SZ"),
        "total": len(projects),
        "summary": summary,
        "projects": projects,
    }


def status_changed(new_data, status_path):
    """Check if status has changed compared to existing status.json."""
    if not os.path.exists(status_path):
        return True

    with open(status_path, "r", encoding="utf-8") as f:
        old_data = json.load(f)

    return old_data.get("projects") != new_data.get("projects")


def main():
    token = os.environ.get("GITHUB_TOKEN")

    if not token:
        print("Warning: GITHUB_TOKEN not set. Unauthenticated API rate limit is 60 req/hour.", file=sys.stderr)
        print("Set GITHUB_TOKEN in .env or environment to get 5000 req/hour.", file=sys.stderr)

    print("Parsing " + README_PATH + "...")
    repos = parse_github_repos(README_PATH)
    print("Found " + str(len(repos)) + " unique GitHub repositories")

    print("Querying GitHub API...")
    projects = check_all_repos(repos, token)

    print("Building status data...")
    status_data = build_status_json(projects)

    print("Summary: " + json.dumps(status_data["summary"]))

    if not status_changed(status_data, STATUS_PATH):
        print("No changes detected. Skipping update.")
        sys.exit(0)

    with open(STATUS_PATH, "w", encoding="utf-8") as f:
        json.dump(status_data, f, indent=2, ensure_ascii=False)

    print("Status written to " + STATUS_PATH)


if __name__ == "__main__":
    main()
