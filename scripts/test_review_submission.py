import unittest
from datetime import datetime, timezone

import review_submission as rs

NOW = datetime(2026, 11, 24, tzinfo=timezone.utc)
LISTED = {"https://github.com/istio/istio": "Service Mesh"}


def facts(**overrides):
    base = {
        "full_name": "acme/widget", "stars": 500, "created_at": "2024-01-01T00:00:00Z",
        "pushed_at": "2026-11-01T00:00:00Z", "license": "Apache-2.0", "fork": False,
        "archived": False, "has_release": True, "description": "", "topics": [],
    }
    base.update(overrides)
    return base


def sub(**overrides):
    base = {"name": "widget", "url": "https://github.com/acme/widget",
            "category": "Service Mesh", "description": "Sidecarless mesh data plane."}
    base.update(overrides)
    return base


def verdict(submission, repo_facts):
    return rs.worst(rs.evaluate(submission, repo_facts, LISTED, NOW))


class ParsingTest(unittest.TestCase):
    def test_issue_form(self):
        body = ("### Project name\n\ntiny\n\n### Repository URL\n\nhttps://github.com/a/tiny\n\n"
                "### Proposed category\n\nDatabases\n\n### One-line description\n\nRuns things.\n")
        self.assertEqual(rs.submission_from_issue(body), {
            "name": "tiny", "url": "https://github.com/a/tiny",
            "category": "Databases", "description": "Runs things."})

    def test_issue_without_form_is_not_a_submission(self):
        self.assertIsNone(rs.submission_from_issue("A link in the list is broken."))

    def test_patch_reports_only_new_entries(self):
        patch = ("@@ -1,3 +1,4 @@\n"
                 " - [a](https://github.com/x/a) - A.\n"
                 "+- [b](https://github.com/x/b) - New.\n"
                 "-- [c](https://github.com/x/c) - Old text.\n"
                 "+- [c](https://github.com/x/c) - Reworded.\n")
        self.assertEqual([e["name"] for e in rs.submissions_from_patch(patch)], ["b"])

    def test_readme_maps_entries_to_categories(self):
        readme = "## Table of Contents\n- [Mesh](#mesh)\n## Mesh\n- [a](https://github.com/x/A/) - A.\n"
        self.assertEqual(rs.parse_readme(readme), {"https://github.com/x/a": "Mesh"})


class EvaluateTest(unittest.TestCase):
    def test_mature_project_is_eligible(self):
        self.assertEqual(verdict(sub(), facts()), "pass")

    def test_hard_requirements_fail(self):
        for repo_facts in (None, facts(fork=True), facts(archived=True), facts(license="BUSL-1.1")):
            self.assertEqual(verdict(sub(), repo_facts), "fail")

    def test_already_listed_fails_even_after_rename(self):
        renamed = sub(url="https://github.com/old/istio")
        self.assertEqual(verdict(renamed, facts(full_name="istio/istio")), "fail")

    def test_immature_project_is_too_early(self):
        for repo_facts in (facts(stars=rs.MIN_STARS - 1), facts(created_at="2026-10-27T00:00:00Z"),
                           facts(has_release=False), facts(pushed_at="2025-01-01T00:00:00Z")):
            self.assertEqual(verdict(sub(), repo_facts), "early")

    def test_content_repositories_need_no_release(self):
        tutorial = sub(category="Tutorials & Learning")
        self.assertEqual(verdict(tutorial, facts(has_release=False)), "pass")

    def test_fixable_problems_need_changes(self):
        self.assertEqual(verdict(sub(name="Acme/Widget"), facts()), "fix")
        self.assertEqual(verdict(sub(description="服务网格"), facts()), "fix")
        self.assertEqual(verdict(sub(), facts(full_name="acme/widget-ng")), "fix")

    def test_unconfirmed_cases_go_to_maintainer(self):
        self.assertEqual(verdict(sub(), facts(license="NOASSERTION")), "manual")
        self.assertEqual(verdict(sub(url="https://gitlab.com/acme/widget"), None), "manual")

    def test_hard_failure_outranks_everything_else(self):
        self.assertEqual(verdict(sub(name="Bad Name"), facts(fork=True, stars=0)), "fail")


class AiReplyTest(unittest.TestCase):
    def test_valid_reply_is_sanitised(self):
        reply = rs.parse_ai_reply('{"relevance": "no", "category_fit": "yes", "notes": "cc @owner [x](http://evil)"}')
        self.assertEqual(reply["relevance"], "no")
        for unsafe in ("@owner", "[", "://"):
            self.assertNotIn(unsafe, reply["notes"])

    def test_malformed_reply_is_rejected(self):
        for content in ("not json", "[]", '{"relevance": "approve", "category_fit": "yes"}', None):
            self.assertIsNone(rs.parse_ai_reply(content))


class RenderTest(unittest.TestCase):
    def test_comment_carries_marker_and_worst_verdict(self):
        reviews = [(sub(), [("pass", "Adoption", "ok")]),
                   (sub(name="other"), [("early", "Adoption", "few stars")])]
        body, level = rs.render("o/r", reviews, is_pr=True)
        self.assertEqual(level, "early")
        self.assertTrue(body.startswith(rs.MARKER))
        self.assertIn("Too early", body)


if __name__ == "__main__":
    unittest.main()
