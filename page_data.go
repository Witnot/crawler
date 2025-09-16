package main

import (
	"net/url"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		// If parsing fails, return empty PageData with URL
		return PageData{URL: pageURL}
	}

	h1 := getH1FromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)

	links, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		links = []string{}
	}

	images, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		images = []string{}
	}

	return PageData{
		URL:            pageURL,
		H1:             h1,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  links,
		ImageURLs:      images,
	}
}
