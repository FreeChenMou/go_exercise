package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var t int

//exercise9.6 矩阵快速幂
func main() {
	now := time.Now()
	var matrix [1000][1000]int
	runtime.GOMAXPROCS(1) //限制线程数量
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			matrix[i][j] = j
		}
	}
	temp := [1000][1000]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			wg.Add(1)
			go func(left, right int) {
				defer wg.Done()
				for i := 0; i < len(matrix[i]); i++ {
					temp[left][right] += matrix[left][i] * matrix[i][right]
				}
			}(i, j)
		}
	}
	wg.Wait()

	fmt.Printf("%s %d", time.Since(now), t)

}
