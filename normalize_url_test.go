package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	testcases := map[string]struct {
		input    []string
		expected []string
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
		},
		"empty string": {
			input: []string{
				"",
			},
			expected: []string{
				"",
			},
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
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			for i, v := range tc.input {
				normalized, err := normalizeURL(v)
				if err == nil && normalized != tc.expected[i] {
					t.Errorf("Expected %v, got %v", tc.expected[i], normalized)
				}
			}
		})
	}
}
