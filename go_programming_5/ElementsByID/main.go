package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

var isExists = false

//exercise5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。
func main() {

	resp, err := http.Get(os.Args[1])
	if err != nil {
		err = fmt.Errorf("http get failed: %v", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %v", err)
		os.Exit(1)
	}
	ElementByID(doc, "img")
}

func ElementByID(doc *html.Node, className string) {
	forEachNode(doc, startElement, endElement, className)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, className string) bool, className string) {
	if pre != nil {
		isExists = pre(n, className)
	}
	for c := n.FirstChild; c != nil && !isExists; c = c.NextSibling {
		forEachNode(c, pre, post, className)
	}

	if post != nil {
		post(n, className)
	}
}

func startElement(n *html.Node, className string) bool {
	if n.Type == html.ElementNode {

		attrs := ""
		for _, a := range n.Attr {
			attrs = attrs + " " + a.Key + "=" + a.Val
		}
		if strings.Contains(attrs, className) {
			fmt.Printf("<%s%s>", n.Data, attrs)
			return true
		}
	}
	return false
}

func endElement(n *html.Node, className string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == className {
				fmt.Printf("</%s>\n", n.Data)
				return true
			}
		}

	}
	return false
}
