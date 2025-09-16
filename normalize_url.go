package main

import (
	"errors"
	"net/url"
)

// normalizeURL takes a URL string and returns a "normalized" version
func normalizeURL(input string) (string, error) {
	if input == "" {
		return "", errors.New("input URL is empty")
	}

	parsed, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	// Remove scheme and return host + path
	normalized := parsed.Host + parsed.Path
	return normalized, nil
}
