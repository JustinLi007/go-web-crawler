package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %v", err)
	}

	if parsedURL.Scheme == "" && parsedURL.Hostname() == "" {
		return "", fmt.Errorf("failed to parse URL: %v", rawURL)
	}

	result := strings.ToLower(fmt.Sprintf(
		"%v%v",
		parsedURL.Hostname(),
		strings.TrimSuffix(parsedURL.Path, "/"),
	))

	return result, nil
}
