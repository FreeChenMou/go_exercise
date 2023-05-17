package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	and := PopCount(35)
	println(and)
	since := time.Since(now)
	fmt.Println(since)

	now = time.Now()
	and = PopCountRange(35)
	println(and)
	since = time.Since(now)
	fmt.Println(since)

	now = time.Now()
	and = PopCount64(35)
	println(and)
	since = time.Since(now)
	fmt.Println(since)

	now = time.Now()
	and = PopCountAnd(35)
	println(and)
	since = time.Since(now)
	fmt.Println(since)
}
