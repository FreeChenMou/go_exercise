package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func compute() {
	now := time.Now()
	defer func() {
		goWordjoin("D:\\project\\GoLand\\src\\go_code\\pride-and-prejudice.txt")
		since := time.Since(now)
		fmt.Println(since)
	}()
}

/*
字符+ 数量大，每次需要生成一个新字符串导致频繁生成和垃圾回收，效率低
*/
func goWordadd(file string) {
	open, err := os.Open(file)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(open)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	var str strings.Builder

	for _, arg := range words[1:] {
		str.WriteString(arg)
		str.WriteString(" ")
	}
}

/*
自我优化版本，通过查看源码stirng源码发现join在源码处是使用Builder库的WriterString进行写入
从而不在生成新的stirng变量和不进行垃圾回收的情况下进行添加，效率提高
*/
func goWordadd2(file string) {
	open, err := os.Open(file)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(open)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	var str strings.Builder
	str.WriteString(words[0])
	for _, arg := range words[1:] {
		str.WriteString(arg)
		str.WriteString(" ")
	}
}

/*
效率高，使用字符封装join，通过writerString搞笑写入拼接
*/
func goWordjoin(file string) {
	open, err := os.Open(file)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(open)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	var _ string
	_ = strings.Join(words[0:], " ")

}
