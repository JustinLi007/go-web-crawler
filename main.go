package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	if numArgs < 2 {
		log.Fatalf("no website provided")
	} else if numArgs > 2 {
		log.Fatalf("too many arguments provided")
	}

	baseURL := os.Args[1]

	const maxConcurrency = 10
	cfg, err := configure(baseURL, maxConcurrency)
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
