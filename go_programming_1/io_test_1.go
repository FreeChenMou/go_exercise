package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test2.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "io_test:%v\n", err)
	}
	open, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "io_rtest:%v\n", err)
	}

	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		file.WriteString("http://" + scanner.Text() + "\n")
	}

}
