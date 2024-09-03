package main

import (
	"fmt"
	"sort"
	"strings"
)

type pageList []page

type page struct {
	key string
	val int
}

// interface for sort
func (p pageList) Len() int {
	return len(p)
}

func (p pageList) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
}

func (p pageList) Less(i, j int) bool {
	if p[i].val == p[j].val {
		return p[i].key < p[j].key
	}
	return p[i].val > p[j].val
}

func printReport(pages map[string]int, baseURL string) {
	const padding = 15
	headerStrArr := []string{}
	for i := 0; i < len(baseURL)+padding; i++ {
		headerStrArr = append(headerStrArr, "=")
	}
	headerString := strings.Join(headerStrArr, "")

	fmt.Println(headerString)
	fmt.Printf("REPORT for %v\n", baseURL)
	fmt.Println(headerString)

	pageList := sortPages(pages)
	//pageList := sortPages2(pages)

	for _, value := range pageList {
		fmt.Printf("Found %v internal links to %v\n", value.val, value.key)
	}
}

func sortPages(pages map[string]int) []page {
	pageList := pageList{}
	for k, v := range pages {
		pageList = append(pageList, page{
			key: k,
			val: v,
		})
	}

	sort.Sort(pageList)
	return pageList
}

func sortPages2(pages map[string]int) []page {
	pageList := []page{}
	for k, v := range pages {
		pageList = append(pageList, page{
			key: k,
			val: v,
		})
	}

	sort.Slice(pageList, func(i, j int) bool {
		if pageList[i].val == pageList[j].val {
			return pageList[i].key < pageList[j].key
		}
		return pageList[i].val > pageList[j].val
	})

	return pageList
}
