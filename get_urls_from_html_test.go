package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	testcases := map[string]struct {
		inputURL  string
		inputBody string
		expected  []string
		error     string
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
			error:    "",
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
			error:    "",
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
			error:    "",
		},
		"invalid hrefs": {
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>`,
			expected: []string{},
			error:    "",
		},
		"invalid base url": {
			inputURL: "://invalidBaseURL",
			inputBody: `
<html>
	<body>
		<a href="/path">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{},
			error:    "failed to parse base URL",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
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
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected urls %v, got %v", tc.expected, actual)
				return
			}
		})
	}
}
