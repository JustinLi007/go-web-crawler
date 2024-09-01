package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, World!")

	htmlBody := `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>
Boot.dev
<a href="https://nesting.in.span.com">
            </span>
		</a>
	</body>
</html>
`

	baseURL := "https://blog.boot.dev"

	allURLs, err := getURLsFromHTML(htmlBody, baseURL)
	if err != nil {
		log.Print(err)
	}

	for _, v := range allURLs {
		fmt.Println(v)
	}
}
