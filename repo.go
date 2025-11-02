package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
	cmd := exec.Command("git", "pull")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error updating repository: %v", err)
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
