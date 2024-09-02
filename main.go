package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	numArgs := len(os.Args)
	if numArgs < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if numArgs > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %v\n", baseURL)

	htmlBody, err := getHTML(baseURL)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(htmlBody)
}
