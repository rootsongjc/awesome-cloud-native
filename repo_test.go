package main

import (
	"bytes"
	"io/ioutil"
	"sort"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

func TestAlpha(t *testing.T) {
	query := startQuery()

	query.Find("body > ul").Each(func(_ int, s *goquery.Selection) {
		testList(t, s)
	})
}

func TestDuplicatedLinks(t *testing.T) {
	query := startQuery()
	links := make(map[string]bool, 0)

	query.Find("body li > a:first-child").Each(func(_ int, s *goquery.Selection) {
		t.Run(s.Text(), func(t *testing.T) {
			href, ok := s.Attr("href")
			if !ok {
				t.Error("expected to have href")
			}

			if links[href] {
				t.Fatalf("duplicated link '%s'", href)
			}

			links[href] = true
		})
	})
}

func TestEnglishDescriptions(t *testing.T) {
	input, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}

	for i, line := range strings.Split(string(input), "\n") {
		if !strings.HasPrefix(line, "- [") {
			continue
		}
		for _, r := range line {
			if isNonLatin(r) {
				t.Errorf("README.md:%d: entry must be in English: %s", i+1, line)
				break
			}
		}
	}
}

func isNonLatin(r rune) bool {
	switch {
	case r >= 0x4E00 && r <= 0x9FFF,
		r >= 0x3400 && r <= 0x4DBF,
		r >= 0x3040 && r <= 0x30FF,
		r >= 0xAC00 && r <= 0xD7AF,
		r >= 0xFF00 && r <= 0xFFEF:
		return true
	}
	return false
}

func TestLinkTextFormat(t *testing.T) {
	input, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}

	for i, line := range strings.Split(string(input), "\n") {
		if !strings.HasPrefix(line, "- [") || !strings.Contains(line, "](http") {
			continue
		}
		end := strings.Index(line, "](")
		name := line[len("- ["):end]
		if strings.ContainsAny(name, "/ ") || strings.ToLower(name) != name {
			t.Errorf("README.md:%d: link text must be lowercase without spaces or slashes: '%s'", i+1, name)
		}
	}
}

func testList(t *testing.T, list *goquery.Selection) {
	list.Find("ul").Each(func(_ int, items *goquery.Selection) {
		testList(t, items)
		items.RemoveFiltered("ul")
	})

	category := list.Prev().Text()

	t.Run(category, func(t *testing.T) {
		checkAlphabeticOrder(t, list)
	})
}

func readme() []byte {
	input, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
	html := append([]byte("<body>"), blackfriday.MarkdownCommon(input)...)
	html = append(html, []byte("</body>")...)
	return html
}

func startQuery() *goquery.Document {
	buf := bytes.NewBuffer(readme())
	query, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		panic(err)
	}

	return query
}

func checkAlphabeticOrder(t *testing.T, s *goquery.Selection) {
	items := s.Find("li > a:first-child").Map(func(_ int, li *goquery.Selection) string {
		return strings.ToLower(li.Text())
	})

	sorted := make([]string, len(items))
	copy(sorted, items)
	sort.Strings(sorted)

	for k, item := range items {
		if item != sorted[k] {
			t.Errorf("expected '%s' but actual is '%s'", sorted[k], item)
		}
	}
	if t.Failed() {
		t.Logf("expected order is:\n%s", strings.Join(sorted, "\n"))
	}
}
