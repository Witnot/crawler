package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	// Limit number of pages
	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	// Acquire concurrency slot
	cfg.concurrencyControl <- struct{}{}
	defer func() { <-cfg.concurrencyControl }() // release slot

	current, err := url.Parse(rawCurrentURL)
	if err != nil || current.Host != cfg.baseURL.Host {
		return
	}

	normalized, err := normalizeURL(current.String())
	if err != nil {
		return
	}

	// Track visits
	cfg.mu.Lock()
	if _, exists := cfg.pages[normalized]; exists {
		cfg.mu.Unlock()
		return
	}
	cfg.pages[normalized] = PageData{} // placeholder
	cfg.mu.Unlock()

	fmt.Printf("Crawling: %s\n", current.String())

	// Fetch HTML
	html, err := getHTML(current.String())
	if err != nil {
		fmt.Printf("failed to fetch %s: %v\n", current.String(), err)
		return
	}

	// Extract page data
	pageData := extractPageData(html, current.String())

	// Store PageData safely
	cfg.mu.Lock()
	cfg.pages[normalized] = pageData
	cfg.mu.Unlock()

	// Extract links and recursively crawl
	links, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		return
	}

	for _, link := range links {
		cfg.mu.Lock()
		if len(cfg.pages) >= cfg.maxPages {
			cfg.mu.Unlock()
			return
		}
		cfg.mu.Unlock()

		cfg.wg.Add(1)
		go func(u string) {
			defer cfg.wg.Done()
			cfg.crawlPage(u)
		}(link)
	}
}
