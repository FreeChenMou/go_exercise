package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

var done = make(chan struct{})
var responses = make(chan string, 3)

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//exercise8.11
func main() {
	Is := true

	for _, url := range os.Args[1:] {
		go mirroredQuery(url)
	}

	for Is {
		select {
		case str := <-responses:
			close(done)
			Is = false
			fmt.Println(str)
		}
	}
}

func mirroredQuery(url string) {
	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	if cancelled() {
		cancel()
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Sprintf("%s request failed.", url)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Sprintf("getting %s: %s", url, resp.Status)
	}

	defer resp.Body.Close()

	responses <- fmt.Sprintf("%s ============== done", url)

}
