package main

import (
	"fmt"
	"os"
	"strconv"
)

//exercise4.3
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := os.Args[1]
	atoi, _ := strconv.Atoi(m)
	reverse(&arr, 0, len(arr)-1)
	reverse(&arr, 0, len(arr)-atoi-1)
	reverse(&arr, len(arr)-atoi, len(arr)-1)

	for _, i2 := range arr {
		fmt.Print(i2)
	}
}

func reverse(arr *[]int, start int, end int) {
	for start < end {
		(*arr)[start], (*arr)[end] = (*arr)[end], (*arr)[start]
		start++
		end--
	}
}
