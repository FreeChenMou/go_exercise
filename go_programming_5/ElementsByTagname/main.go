package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

//exercise5.17 变长函数ElementsByTagname
func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get err:%v", err)
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http parse err:%v", err)
	}
	var byTagname []*html.Node
	byTagname = ElementsByTagname(byTagname, node, "a", "img")
	for _, h := range byTagname {
		fmt.Println(h.Data)
	}
}

func ElementsByTagname(node []*html.Node, doc *html.Node, name ...string) []*html.Node {
	if doc.Type == html.ElementNode {
		for _, s := range name {
			if doc.Data == s {
				node = append(node, doc)
			}
		}
	}
	for n := doc.FirstChild; n != nil; n = n.NextSibling {
		node = ElementsByTagname(node, n, name...)
	}
	return node

}
