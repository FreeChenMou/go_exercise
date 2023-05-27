package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

var count = make(chan int)
var record int

//exercise9.4
func main() {
	getMemSys := func() uint64 {
		mem := runtime.MemStats{}
		runtime.ReadMemStats(&mem)
		return mem.Sys
	}
	before := getMemSys()
	go func() {
		for {
			go func() {
				i := <-count
				record++
				count <- i
			}()
		}
	}()

	count <- 0
	now := time.Now()

	for {
		memAlloc := getMemSys() - before
		if memAlloc > uint64(2*math.Pow(2, 30)) {
			fmt.Printf("time:%s level :%d", time.Since(now), record)
			break
		}
	}

}
