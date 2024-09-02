package main

import (
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, err
	}

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
				if v.Key == "href" {
					href, err := url.Parse(v.Val)
					if err != nil {
						continue
					}
					resolvedURL := baseURL.ResolveReference(href)
					result = append(result, resolvedURL.String())
				}
			}
		}

		traverse(node.FirstChild)
		traverse(node.NextSibling)
	}

	traverse(doc.FirstChild)

	return result, nil
}
