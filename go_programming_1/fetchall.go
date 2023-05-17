// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fetchall2()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		//return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.4fs  %7d  %s\n", secs, nbytes, url)
}

//exercise1.10 内容输出到文件，查看多次获取资源时间变化以此来查看缓存情况
func fetchall1() {
	open, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall1:%v\n", err)
	}
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		//fmt.Println(<-ch) // receive from channel ch
		open.WriteString(<-ch)
	}
	open.WriteString(fmt.Sprintf("%.4fs elapsed\n", time.Since(start).Seconds()))
}

//exercise1.11 多并发尝试
func fetchall2() {
	start := time.Now()
	ch := make(chan string)
	open, err := os.Open("test2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall2:%v\n", err)
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		go fetch(scanner.Text(), ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.4fs elapsed\n", time.Since(start).Seconds())
}
