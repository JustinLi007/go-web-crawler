package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		log.Printf("failed to parse html body: %v", err)
	}

	result := []string{}

	var traverse func(node *html.Node)
	traverse = func(node *html.Node) {
		if node == nil {
			return
		}

		if node.Type == html.ElementNode && node.Data == "a" {
			for _, v := range node.Attr {
				result = append(result, v.Val)
			}
		}

		traverse(node.FirstChild)
		traverse(node.NextSibling)
	}

	traverse(doc.FirstChild)

	for i, v := range result {
		parseURL, err := url.Parse(v)
		if err != nil {
			return []string{}, nil
		}
		if !parseURL.IsAbs() {
			result[i] = fmt.Sprintf("%v%v", rawBaseURL, v)
		}
	}

	return result, nil
}
