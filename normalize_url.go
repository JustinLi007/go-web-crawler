package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Scheme == "" && parsedURL.Hostname() == "" {
		return "", errors.New(fmt.Sprintf("url %v have no scheme or domain", rawURL))
	}

	result := fmt.Sprintf("%v%v", parsedURL.Hostname(), strings.TrimSuffix(parsedURL.Path, "/"))

	return result, nil
}
