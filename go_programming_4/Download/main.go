package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type picture struct {
	Poster string
}

var BaseUrl = "https://omdbapi.com/"
var ApiKey = "4718a23a"

func main() {
	search()
}

//exercise4.13 根据电影名或描述下载海报
func search() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 1
	for scanner.Scan() {
		url := fmt.Sprintf(BaseUrl+"?t=%s&apikey="+ApiKey, scanner.Text())
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "search err:%v", err)
		}
		var result picture
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			fmt.Fprintf(os.Stderr, "json decoder:%v \n", err)
		}
		response, err := http.Get(result.Poster)
		bytes, err := ioutil.ReadAll(response.Body)
		png := fmt.Sprintf("out%d.png", count)
		count++
		if err := ioutil.WriteFile(png, bytes, 0666); err != nil {
			panic(err)
		}
	}

}
