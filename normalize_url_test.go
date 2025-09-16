package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http scheme",
			inputURL: "http://example.com/test",
			expected: "example.com/test",
		},
		{
			name:     "no scheme",
			inputURL: "example.com/noscheme",
			expected: "example.com/noscheme",
		},
		{
			name:     "empty input",
			inputURL: "",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if tc.inputURL == "" && err == nil {
				t.Errorf("Test %v - '%s' FAIL: expected error for empty input", i, tc.name)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
