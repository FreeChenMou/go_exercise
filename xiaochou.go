package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	record := make(map[string]int)
	mutex := sync.RWMutex{}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v:", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		go func() {
			mutex.Lock()
			record[scanner.Text()]++
			mutex.Unlock()
		}()
	}

}
