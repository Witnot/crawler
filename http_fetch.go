package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// getHTML fetches the HTML content of the provided URL
func getHTML(rawURL string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}

	// Set User-Agent to avoid blocks
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check for HTTP error codes (400+)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}

	// Check Content-Type
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" || !containsHTML(contentType) {
		return "", errors.New("response is not HTML")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// Helper: check if Content-Type contains "text/html"
func containsHTML(contentType string) bool {
	return len(contentType) >= 9 && contentType[:9] == "text/html"
}
