package main

import "time"

var count = make(chan int)
var record int

//exercise9.5
func main() {

	go func() {
		for {
			<-count
			record++
			count <- record
		}
	}()

	go func() {
		for {
			<-count
			record++
			count <- record
		}
	}()
	count <- record
	for {
		ticker := time.NewTicker(1 * time.Second)
		select {
		case <-ticker.C:
			println(record)
		}
	}
}
