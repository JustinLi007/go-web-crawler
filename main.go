package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	numArgs := len(os.Args)
	if numArgs < 2 {
		fmt.Printf("no website provided")
		os.Exit(1)
	} else if numArgs > 2 {
		fmt.Printf("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %v\n", baseURL)

	htmlBody, err := getHTML(baseURL)
	if err != nil {
		log.Printf("failed to get html from %v: %v", baseURL, err)
	}

	fmt.Println(htmlBody)
}
