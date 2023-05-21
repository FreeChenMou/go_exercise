package main

import "fmt"

//exercise5.19 抛出异常并接收处理返回
func main() {

	code := excute(0)
	fmt.Println(code)
}

func excute(x int) (code int) {
	defer func() {
		if err := recover(); err == -1 {
			code = err.(int)
		}
	}()
	panic(-1)
}
