package main

import (
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	testcases := map[string]struct {
		inputURL  string
		inputBody string
		expected  []string
	}{
		"absolute and relative urls": {
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		"nested absolute and relative urls": {
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>
            Boot.dev
            <a href="http://nested.in.span.com">
            <a href="/nested/relative">
            </span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one", "http://nested.in.span.com", "https://blog.boot.dev/nested/relative"},
		},
		"no href": {
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a >
			<span>Boot.dev</span>
		</a>
		<a >
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			actual, _ := getURLsFromHTML(tc.inputBody, tc.inputURL)
			for i, v := range actual {
				if tc.expected[i] != v {
					t.Errorf("Expected %v, got %v", tc.expected[i], v)
				}
			}
		})
	}
}
