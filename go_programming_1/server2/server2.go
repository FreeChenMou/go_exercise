package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var p sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	p.Lock()
	count++
	p.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request) {
	p.Lock()
	fmt.Fprintf(w, "当前访问次数:%d", count)
	count++
	p.Unlock()
}
