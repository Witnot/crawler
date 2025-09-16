package main

import (
	"encoding/csv"
	"os"
	"strings"
)

// writeCSVReport writes the crawled pages into a CSV file
func writeCSVReport(pages map[string]PageData, filename string) error {
	// Create or overwrite the CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	header := []string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write each PageData as a row
	for url, page := range pages {
		row := []string{
			url,
			page.H1,
			page.FirstParagraph,
			strings.Join(page.OutgoingLinks, ";"),
			strings.Join(page.ImageURLs, ";"),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
