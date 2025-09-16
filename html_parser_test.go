package main

import "testing"

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic h1",
			input:    "<html><body><h1>Test Title</h1></body></html>",
			expected: "Test Title",
		},
		{
			name:     "no h1",
			input:    "<html><body><p>No H1 here</p></body></html>",
			expected: "",
		},
		{
			name:     "multiple h1",
			input:    "<h1>First</h1><h1>Second</h1>",
			expected: "First",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "first p in main",
			input: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
			</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "no main, use first p",
			input: `<html><body>
				<p>First paragraph.</p>
				<p>Second paragraph.</p>
			</body></html>`,
			expected: "First paragraph.",
		},
		{
			name:     "no p at all",
			input:    "<html><body><h1>Only H1</h1></body></html>",
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}
