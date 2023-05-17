package main

import "fmt"

//exercise4.4 元素旋转
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	strat, end := 0, len(arr)-1
	for strat < end {
		arr[strat], arr[end] = arr[end], arr[strat]
		strat++
		end--
	}

	for _, i2 := range arr {
		fmt.Print(i2)
	}

}
