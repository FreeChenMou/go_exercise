package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type skcd struct {
	Link string
	Alt  string
}

var mutex sync.RWMutex
var wg sync.WaitGroup

//exercise4.12 构建离线索引进行查找
func main() {
	record := make(map[string]skcd)

	for i := 1; i < 2000; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		go handle(url, record)
	}

	wg.Wait()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		atoi, _ := strconv.Atoi(scanner.Text())
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", atoi)
		fmt.Println(record[url].Alt)
	}

}

func handle(url string, record map[string]skcd) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get err:%v", err)
	}
	var result skcd
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decoder err:%v\n", err)
		resp.Body.Close()
	}
	resp.Body.Close()
	mutex.Lock()
	record[url] = result
	mutex.Unlock()

}
