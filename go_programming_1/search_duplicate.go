package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	exercise_dup2()
}

func dup1() {
	consts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	i, _ := strconv.Atoi(input.Text())
	for j := 0; j < i; j++ {
		input.Scan()
		consts[input.Text()]++
	}

	for keys, value := range consts {
		if value > 1 {
			fmt.Printf("%d %s\n", value, keys)
		}
	}
}

//从文件列表或者标准输入读取
func dup2() {
	consts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, consts)
	} else {
		for _, org := range files {
			file, err := os.Open(org)
			if nil != err {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				consts[scanner.Text()]++
			}
			file.Close()
		}
	}
	for key, value := range consts {
		if value > 1 {
			fmt.Printf("words:%s count:%d\n", key, value)
		}
	}
}

func countLines(stdin *os.File, consts map[string]int) {
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		consts[scanner.Text()]++
	}
}

//单独读取文件，一次读取所有数据
func dup3() {
	count := make(map[string]int)
	for _, org := range os.Args[1:] {
		file, err := ioutil.ReadFile(org)
		if nil != err {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, text := range strings.Split(string(file), "\n") {
			count[text]++
		}
	}
	for key, value := range count {
		if value > 1 {
			fmt.Printf("count:%d words:%s \n", value, key)
		}
	}
}

//输出出现重复行的文件的名称
func exercise_dup2() {
	for _, arg := range os.Args[1:] {
		count := make(map[string]int)
		file, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			count[scanner.Text()]++
		}
		file.Close()
		for _, value := range count {
			if value > 1 {
				fmt.Printf("duplicate file:%s\n", file.Name())
				break
			}
		}
	}
}
