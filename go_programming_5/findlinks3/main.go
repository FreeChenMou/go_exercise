package main

import (
	"go_code/go_exercise/go_programming_5/links"
	"strings"

	"log"
	"os"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if strings.HasPrefix(item, "http://www.baidu.com") { //只有当前域名才能进入访问
					worklist = append(worklist, f(item)...)
				}
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
