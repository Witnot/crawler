package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// getURLsFromHTML extracts all <a href> links from the HTML.
// It converts relative URLs to absolute URLs based on baseURL.
func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var urls []string
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists || href == "" {
			return
		}
		absURL, err := baseURL.Parse(href)
		if err != nil {
			return
		}
		urls = append(urls, absURL.String())
	})

	return urls, nil
}

// getImagesFromHTML extracts all <img src> URLs from the HTML.
// Converts relative URLs to absolute URLs based on baseURL.
func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var imgs []string
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists || src == "" {
			return
		}
		absURL, err := baseURL.Parse(src)
		if err != nil {
			return
		}
		imgs = append(imgs, absURL.String())
	})

	return imgs, nil
}
