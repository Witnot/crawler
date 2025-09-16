package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	baseURLStr := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil || maxConcurrency < 1 {
		fmt.Println("maxConcurrency must be a positive integer")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil || maxPages < 1 {
		fmt.Println("maxPages must be a positive integer")
		os.Exit(1)
	}

	base, err := url.Parse(baseURLStr)
	if err != nil {
		fmt.Printf("invalid URL: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", baseURLStr)

	cfg := &config{
		pages:              make(map[string]PageData),
		baseURL:            base,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.wg.Add(1)
	go func() {
		defer cfg.wg.Done()
		cfg.crawlPage(baseURLStr)
	}()

	cfg.wg.Wait()
	if err := writeCSVReport(cfg.pages, "report.csv"); err != nil {
		fmt.Printf("failed to write CSV report: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("CSV report written to report.csv")
	fmt.Println("\nCrawled pages:")
	cfg.mu.Lock()
	for url, data := range cfg.pages {
		fmt.Printf("%s -> H1: %q, Paragraph: %q, Links: %d, Images: %d\n",
			url, data.H1, data.FirstParagraph, len(data.OutgoingLinks), len(data.ImageURLs))
	}
	cfg.mu.Unlock()
}
