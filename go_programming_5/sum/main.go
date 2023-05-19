// The sum program demonstrates a variadic function.
package main

import "fmt"

//!+
func max(max ...int) int {
	m := 0x8000000
	if len(max) == 0 {
		fmt.Println("参数个数至少大于等于1")
		return -1
	}
	for _, i := range max {
		if m < i {
			m = i
		}
	}
	return m
}

func min(min ...int) int {
	m := 0x7fffffff
	if len(min) == 0 {
		fmt.Println("参数个数至少大于等于1")
		return -1
	}
	for _, i := range min {
		if m > i {
			m = i
		}
	}
	return m
}

//!-
//exercise5.15 变长max和变长min
func main() {
	fmt.Println(max())

	fmt.Println(min(1, 2, 3, 5, 0))
}
