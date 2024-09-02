package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type PageList []Page

type Page struct {
	key string
	val int
}

func (p PageList) Len() int {
	return len(p)
}

func (p PageList) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
}

func (p PageList) Less(i, j int) bool {
	return p[i].val > p[j].val
}

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

	pageList := PageList{}
	for k, v := range cfg.pages {
		pageList = append(pageList, Page{
			key: k,
			val: v,
		})
	}

	sort.Sort(pageList)

	for _, value := range pageList {
		fmt.Printf("%v: %v\n", value.key, value.val)
	}
}
