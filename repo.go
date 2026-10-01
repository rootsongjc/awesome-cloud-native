package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	gfm "github.com/shurcooL/github_flavored_markdown"
)

const (
	readmePath = "./README.md"
	tplPath    = "tmpl/tmpl.html"
	idxPath    = "tmpl/index.html"
)

type content struct {
	Body string
}

func updateRepo() {
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := branchCmd.Output()
	if err != nil {
		log.Printf("Skipping git pull (unable to detect branch): %v", err)
		return
	}

	currentBranch := strings.TrimSpace(string(out))
	if currentBranch == "HEAD" || currentBranch == "" {
		log.Println("Skipping git pull: detached HEAD detected")
		return
	}

	pullCmd := exec.Command("git", "pull", "--ff-only")
	pullCmd.Stdout = os.Stdout
	pullCmd.Stderr = os.Stderr
	if err := pullCmd.Run(); err != nil {
		log.Printf("Continuing without git pull (failed to update repository): %v", err)
	}
}

func readMarkdownFile() []byte {
	input, err := ioutil.ReadFile(readmePath)
	if err != nil {
		log.Fatalf("Error reading README.md: %v", err)
	}
	return input
}

func generateHTML(input []byte) {
	body := string(gfm.Markdown(input))
	body = fixAnchors(body)
	c := &content{Body: body}

	t, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	f, err := os.Create(idxPath)
	if err != nil {
		log.Fatalf("Error creating index.html: %v", err)
	}
	defer f.Close()

	if err := t.Execute(f, c); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	updateRepo()
	markdown := readMarkdownFile()
	generateHTML(markdown)
	log.Println("Successfully generated index.html")
}

var (
	// headingRe matches a heading element emitted by the GFM renderer,
	// capturing the level and the heading's inner HTML.
	headingRe = regexp.MustCompile(`(?s)<h([1-6])><a name="[^"]*" class="anchor" href="[^"]*" rel="nofollow" aria-hidden="true"><span class="octicon octicon-link"></span></a>(.*?)</h[1-6]>`)
	tagRe     = regexp.MustCompile(`(?s)<[^>]*>`)
)

// githubSlug returns the heading anchor GitHub generates for a heading title:
// lowercased, punctuation removed (each removed run still leaves its surrounding
// spaces behind), and spaces replaced with hyphens without collapsing repeats.
func githubSlug(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '_', r == '-', r == ' ':
			b.WriteRune(r)
		}
	}
	return strings.ReplaceAll(b.String(), " ", "-")
}

// fixAnchors rewrites heading anchor names to match GitHub's slug algorithm so
// that in-page links (e.g. the Table of Contents) resolve identically on GitHub
// and on the locally generated static site.
func fixAnchors(body string) string {
	return headingRe.ReplaceAllStringFunc(body, func(m string) string {
		sm := headingRe.FindStringSubmatch(m)
		if len(sm) < 3 {
			return m
		}
		level := sm[1]
		inner := sm[2]
		title := html.UnescapeString(tagRe.ReplaceAllString(inner, ""))
		slug := githubSlug(title)
		return fmt.Sprintf(`<h%s><a name="%s" class="anchor" href="#%s" rel="nofollow" aria-hidden="true"><span class="octicon octicon-link"></span></a>%s</h%s>`,
			level, slug, slug, inner, level)
	})
}
