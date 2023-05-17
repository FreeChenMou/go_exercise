package main

import (
	"bytes"
	"strings"
)

//exercise3.11 正确处理浮点数
func main() {
	var buf bytes.Buffer
	var number string = "1234.6789"
	lastIndex := strings.LastIndex(number, ".")
	newNumber := number[0:lastIndex]

	newLength := len(newNumber)
	index := newLength % 3
	buf.WriteByte('[')
	buf.WriteString(newNumber[:index])
	if index > 0 && index < newLength {
		buf.WriteByte(',')
	}
	for i := index; i < newLength; i++ {
		buf.WriteString(newNumber[i : i+3])
		if i+3 == newLength {
			break
		}
		buf.WriteByte(',')
	}
	buf.WriteString(number[lastIndex:])
	buf.WriteByte(']')
	s := buf.String()
	println(s)
}
