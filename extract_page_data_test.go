package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		pageURL  string
		expected PageData
	}{
		{
			name:    "basic page",
			pageURL: "https://blog.boot.dev",
			html: `<html><body>
				<h1>Test Title</h1>
				<p>This is the first paragraph.</p>
				<a href="/link1">Link 1</a>
				<img src="/image1.jpg" alt="Image 1">
			</body></html>`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
				ImageURLs:      []string{"https://blog.boot.dev/image1.jpg"},
			},
		},
		{
			name:    "no h1 or p",
			pageURL: "https://example.com",
			html:    `<html><body><a href="/link1">Link 1</a><img src="/img.png"></body></html>`,
			expected: PageData{
				URL:            "https://example.com",
				H1:             "",
				FirstParagraph: "",
				OutgoingLinks:  []string{"https://example.com/link1"},
				ImageURLs:      []string{"https://example.com/img.png"},
			},
		},
		{
			name:    "multiple links and images",
			pageURL: "https://example.com",
			html: `<html><body>
				<h1>Title</h1>
				<p>Para</p>
				<a href="/one">One</a>
				<a href="https://example.com/two">Two</a>
				<img src="/img1.png">
				<img src="https://example.com/img2.png">
			</body></html>`,
			expected: PageData{
				URL:            "https://example.com",
				H1:             "Title",
				FirstParagraph: "Para",
				OutgoingLinks:  []string{"https://example.com/one", "https://example.com/two"},
				ImageURLs:      []string{"https://example.com/img1.png", "https://example.com/img2.png"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.html, tc.pageURL)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %+v, got %+v", tc.expected, actual)
			}
		})
	}
}
