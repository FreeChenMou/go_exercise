package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	var while_str, while_sep string
	length := len(os.Args)
	var index int
	for index < length {
		while_str += while_sep + os.Args[index]
		print(os.Args[index])
		println(index)
		while_sep = " "
		index++
	}

	var str, sep string
	//for循环遍历命令行的两种方法 index and range
	for i, arg := range os.Args {
		str += sep + arg
		print(arg)
		println(i)
		sep = " "
	}

	var str2, sep2 string
	for i := 0; i < len(os.Args); i++ {
		str2 += sep2 + os.Args[i]
		print(os.Args[i])
		println(i)
		sep2 = " "
	}
}
