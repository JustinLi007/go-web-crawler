package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.isPageLimit() {
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("error - crawlPage: failed to parse URL %v: %v", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	normCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("error - normalizeURL: %v", err)
		return
	}

	if !cfg.addPageVisit(normCurrent) {
		return
	}

	fmt.Printf("crawling %v\n", rawCurrentURL)
	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("error - getHTML: %v", err)
		return
	}

	urls, err := getURLsFromHTML(currentHTML, cfg.baseURL)
	if err != nil {
		log.Printf("error - getURLsFromHTML: %v", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
