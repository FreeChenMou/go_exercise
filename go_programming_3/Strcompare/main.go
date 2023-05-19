package main

import "fmt"

//exercise3.12  判断两个字符串是否异构
func main() {
	var str1, str2 string = "abcdf", "fdcba"

	record := make(map[byte]int)
	var temp bool = true
	if len(str1) != len(str2) {
		fmt.Println("否")
		return
	}
	for i := 0; i < len(str1); i++ {
		record[str1[i]]++
		record[str2[i]]--
	}

	for _, i := range record {
		if i > 0 {
			temp = false
		}
	}
	println(temp)
}
