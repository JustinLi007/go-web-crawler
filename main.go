package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numArgs := len(os.Args)
	if numArgs < 4 {
		log.Printf("no website provided")
		fmt.Printf("usage: crawler <url> <maxConcurrency> <maxPages>")
		return
	} else if numArgs > 4 {
		log.Fatalf("too many arguments provided")
	}

	baseURL := os.Args[1]
	maxConcurrencyStr := os.Args[2]
	maxPagesStr := os.Args[3]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyStr)
	if err != nil {
		log.Printf("error - maxConcurrency: %v", err)
	}
	maxPages, err := strconv.Atoi(maxPagesStr)
	if err != nil {
		log.Printf("error - maxPages: %v", err)
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPages)
	if err != nil {
		log.Fatalf("error - configure: %v", err)
	}

	fmt.Printf("starting crawl of: %v\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)

	cfg.wg.Wait()

	printReport(cfg.pages, baseURL)
}
