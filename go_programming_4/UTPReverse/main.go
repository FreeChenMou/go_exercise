package main

import (
	"fmt"
	"unicode/utf8"
)

//exercise4.7 翻转UTF-8字符串
func main() {
	var str = "阪大上年明epoh"
	bytes := []byte(str)

	reverse(bytes)

	fmt.Println(string(bytes))

}

func reverse(bytes []byte) {
	length := len(bytes)

	for length > 0 {
		inString, size := utf8.DecodeRuneInString(string(bytes))

		copy(bytes[0:], bytes[0+size:length])
		copy(bytes[length-size:length], []byte(string(inString)))
		length -= size
	}
}
