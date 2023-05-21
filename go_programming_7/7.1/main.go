package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (c *ByteCounter) CountWord(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*c++
	}
	return len(scanner.Bytes()), nil
}

func (c *ByteCounter) CountLine(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		*c++
	}
	return len(scanner.Bytes()), nil
}

//exercise7.1 计算行与列
func main() {
	var c ByteCounter
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "open file err:%v", err)
	}
	c.CountLine(file)
	defer file.Close()
	println(c)
}
