package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	testcases := map[string]struct {
		input    []string
		expected []string
		error    string
	}{
		"normal urls": {
			input: []string{
				"https://blog.boot.dev/path/",
				"https://blog.boot.dev/path",
				"http://blog.boot.dev/path/",
				"http://blog.boot.dev/path",
			},
			expected: []string{
				"blog.boot.dev/path",
				"blog.boot.dev/path",
				"blog.boot.dev/path",
				"blog.boot.dev/path",
			},
			error: "",
		},
		"mixed cases": {
			input: []string{
				"https://BLOG.boot.dev/PATH",
				"http://BLOG.boot.dev/path/",
			},
			expected: []string{
				"blog.boot.dev/path",
				"blog.boot.dev/path",
			},
			error: "",
		},
		"empty string": {
			input: []string{
				"",
			},
			expected: []string{
				"",
			},
			error: "failed to parse URL",
		},
		"bad urls": {
			input: []string{
				"blog",
				":\\path",
			},
			expected: []string{
				"",
				"",
			},
			error: "failed to parse URL",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			for i, v := range tc.input {
				actual, err := normalizeURL(v)
				if err != nil && !strings.Contains(err.Error(), tc.error) {
					t.Errorf("Unexpected error: %v", err)
					return
				}
				if err != nil && tc.error == "" {
					t.Errorf("Unexpected error: %v", err)
					return
				}
				if err == nil && tc.error != "" {
					t.Errorf("Expected error %v, got none", tc.error)
					return
				}
				if actual != tc.expected[i] {
					t.Errorf("Expected %v, got %v", tc.expected[i], actual)
					return
				}
			}
		})
	}
}
