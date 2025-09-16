package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// getH1FromHTML returns the text content of the first <h1> tag.
// Returns empty string if no <h1> exists.
func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	h1 := doc.Find("h1").First()
	if h1.Length() == 0 {
		return ""
	}

	return strings.TrimSpace(h1.Text())
}

// getFirstParagraphFromHTML returns the text content of the first <p> tag.
// If <main> exists, finds the first <p> within <main>. Otherwise, first <p> in the document.
func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	var p *goquery.Selection
	main := doc.Find("main").First()
	if main.Length() > 0 {
		p = main.Find("p").First()
	} else {
		p = doc.Find("p").First()
	}

	if p.Length() == 0 {
		return ""
	}

	return strings.TrimSpace(p.Text())
}
