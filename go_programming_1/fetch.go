// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	fetch3()
}

//exercise1.7 copy输出流
func fetch1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		dst := make([]byte, 1024*4)
		/*1.通过声明一个writer流copy数据
				var write io.Writer
		        io.Copy(write,resp.Body)
		*/
		/*2.通过io工具使用io.Discard copy数据*/
		_, err = io.Copy(ioutil.Discard, resp.Body)

		/*3.系统调用syscall.Stdout输出流
		io.Copy(os.Stdout, resp.Body)
		*/
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		ioutil.Discard.Write(dst)
		//ioutil.Discard.Write(dst)
		fmt.Printf("%s", dst)
	}
}

//exercise1.8 弥补前缀
func fetch2() {
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "http")
		if !hasPrefix {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

//exercise1.9 http状态码
func fetch3() {
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "http")
		if !hasPrefix {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("状态码:%s \n", resp.Status)
		fmt.Printf("%s", b)
	}
}
