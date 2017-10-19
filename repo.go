package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	gfm "github.com/shurcooL/github_flavored_markdown"
)

// memory usage optimizations
const (
	emtyStr  = ""
	git      = "git"
	checkout = "checkout"
	force    = "-f"
	pull     = "pull"

	// options
	readmePath = "./README.md"
	tplPath    = "tmpl/tmpl.html"
	idxPath    = "tmpl/index.html"
)

var (
	doneResp = []byte("Done!\n")
)

type content struct {
	Body string
}

func generateHTML() {
	// Update repo
	exec.Command(git, checkout, force).Output()
	exec.Command(git, pull).Output()

	input, _ := ioutil.ReadFile(readmePath)
	body := string(gfm.Markdown(input))
	c := &content{Body: body}

	t := template.Must(template.ParseFiles(tplPath))
	f, _ := os.Create(idxPath)
	t.Execute(f, c)
}

func main() {
	generateHTML()
}
