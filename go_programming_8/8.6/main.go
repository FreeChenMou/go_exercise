package main

import (
	"fmt"
	"log"
	"os"
)

var token = make(chan struct{}, 20)

const depth = 3

var count int64 = 0

func crawl(url string) []string {
	fmt.Println(url)
	if count >= depth {
		return nil
	}
	count++
	token <- struct{}{}
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	<-token
	return list
}

//!+
func main() {
	worklist := make(chan []string) // lists of URLs, may have duplicates

	var n int
	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()
	n++
	// Create 20 crawler goroutines to fetch each unseen link.

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) { worklist <- crawl(link) }(link)
			}
		}
	}
}
