package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func httpGet(url string) ([]string, error) {
	//resp, err := http.Get(url)
	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	if cancelled() {
		cancel()
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	resp.Body.Close()
	fmt.Printf("%s ============== done \n", url)
	return ExtractWithoutRequest(url, bodyStr, resp)
}

func crawl(url string) []string {
	list, err := httpGet(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// 练习 8.10： HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消。
var done = make(chan struct{})

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	go func() { worklist <- os.Args[1:] }()
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()
	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
