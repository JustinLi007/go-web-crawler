package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return "", errors.New(resp.Status)
	}
	if contentType := resp.Header.Get("Content-Type"); !strings.HasPrefix(contentType, "text/html") {
		return "", errors.New(fmt.Sprintf("expected content type %v, got %v", "text/html", contentType))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	htmlBody := string(data)

	return htmlBody, nil
}
