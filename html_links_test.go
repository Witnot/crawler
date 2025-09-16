package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "absolute URL",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="https://blog.boot.dev">Boot.dev</a></body></html>`,
			expected:  []string{"https://blog.boot.dev"},
		},
		{
			name:      "relative URL",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="/path">Relative</a></body></html>`,
			expected:  []string{"https://blog.boot.dev/path"},
		},
		{
			name:      "multiple links",
			inputURL:  "https://example.com",
			inputBody: `<a href="/one">One</a><a href="https://example.com/two">Two</a>`,
			expected:  []string{"https://example.com/one", "https://example.com/two"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Fatalf("couldn't parse input URL: %v", err)
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "relative image",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name:      "absolute image",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="https://blog.boot.dev/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name:      "multiple images with missing src",
			inputURL:  "https://example.com",
			inputBody: `<img src="/one.png"><img><img src="https://example.com/two.png">`,
			expected:  []string{"https://example.com/one.png", "https://example.com/two.png"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Fatalf("couldn't parse input URL: %v", err)
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
