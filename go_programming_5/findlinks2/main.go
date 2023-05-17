package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	links, err := findlinks(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "find err:%v", err)
	}
	for _, link := range links {
		fmt.Println(link)
	}

	words, images, err := CountWordAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "count err:%v", err)
	}
	fmt.Println(words, images)

}

func findlinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get err:%v", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get %s err:%v", url, resp.Status)
	}
	parse, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parse html err:%v", err)
	}

	return visit(nil, parse), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}

		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func CountWordAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("http get err:%v", err)
		return
	}

	parse, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parse err:%v", err)
		return
	}
	words, images = countWordAndImages(parse)

	return
}

//exercise5.5 计算单词和图片的数量
func countWordAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		word, image := countWordAndImages(c)
		words += word
		images += image
	}
	return words, images
}
