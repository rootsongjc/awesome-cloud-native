package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
