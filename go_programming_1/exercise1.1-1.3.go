package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	exercise_1()
	exercise_2()
	exercise_3()

}

func exercise_1() {
	//exercise 1.1
	println(os.Args[0])
}

func exercise_2() {
	//exercise 1.2
	var str, sep2 string
	for arg := range os.Args {
		str += sep2 + os.Args[arg]
		print(arg)
		println(os.Args[arg])
		sep2 = " "
	}
}

func exercise_3() {
	//exercise 1.3
	//compute()
	now2 := time.Now()
	goWordadd("D:\\project\\GoLand\\src\\go_code\\pride-and-prejudice.txt")
	since2 := time.Since(now2)
	fmt.Println(since2)
	now := time.Now()
	goWordadd2("D:\\project\\GoLand\\src\\go_code\\pride-and-prejudice.txt")
	since := time.Since(now)
	fmt.Println(since)
}
