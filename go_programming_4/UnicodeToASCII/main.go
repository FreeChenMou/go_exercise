package main

import (
	"fmt"
	"unicode"
)

//exercise4.6 去除相邻空白字符
func main() {
	var str = "worinm  cs   cnm  "
	bytes := []byte(str)
	for i, j := 0, 1; j < len(bytes); i, j = i+1, j+1 {
		if unicode.IsSpace(rune(bytes[i])) && unicode.IsSpace(rune(bytes[j])) {
			copy(bytes[i:], bytes[j:])
			bytes = bytes[:len(bytes)-1]
			i--
			j--
		}
	}
	fmt.Println(string(bytes))
}
