package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

//练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，
func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("get url:%s failed: %s", url, err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		elementPrefix := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
		attrs := ""
		for _, a := range n.Attr {
			attrVal := a.Val
			if a.Key == "href" || a.Key == "src" {
				attrVal = `"` + attrVal + `"`
			}
			attrs = attrs + " " + a.Key + "=" + attrVal
		}

		if n.FirstChild == nil {
			fmt.Printf("%s /> \n", elementPrefix+attrs)
		} else {
			fmt.Printf("%s> \n", elementPrefix+attrs)
		}
		depth++
		if n.FirstChild != nil && (n.FirstChild.Type == html.CommentNode || n.FirstChild.Type == html.TextNode) {
			if len(strings.TrimSpace(n.FirstChild.Data)) > 0 {
				fmt.Printf("%*s<%s>\n", depth*2+2, "", strings.TrimSpace(n.FirstChild.Data))
			}
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
