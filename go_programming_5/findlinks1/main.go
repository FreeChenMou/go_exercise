// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	var text []string
	for _, text := range visit4(text, doc) {
		fmt.Println(text)
	}

	//exercise5.2
	//record := make(map[string]int)
	//for key, value := range visit(record, doc) {
	//	fmt.Printf("element:%s count:%d\n", key, value)
	//}

}

//!-main

//!+visit
//exercise5.2 循环统计元素
func visit(record map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		record[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		record = visit(record, c)
	}
	return record
}

//!+visit
//exercise5.1 visit appends to links each link found in n and returns the result.递归
func visit2(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	links = visit2(links, n.FirstChild)
	links = visit2(links, n.NextSibling)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	return links
}

//exercise5.3 统计所有文本节点的内容
func visit3(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" && c.Data == "style" {
			continue
		}
		texts = visit3(texts, c)
	}
	return texts
}

//exercise5.4 其他种类的链接
func visit4(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode &&
		(n.Data == "img" || n.Data == "a" || n.Data == "link" || n.Data == "scripts") {
		for _, key := range n.Attr {
			if key.Key == "href" {
				links = append(links, key.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit4(links, c)
	}
	return links
}

//!-visit

/*
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
*/
