package main

import "bytes"

//exercise3.10 非递归comma
func main() {
	var buf bytes.Buffer
	var number string = "123"
	buf.WriteByte('[')
	index := len(number) % 3
	buf.WriteString(number[0:index])
	if index > 0 && index < len(number) {
		buf.WriteByte(',')
	}
	for i := index; i < len(number); i += 3 {

		buf.WriteString(number[i : i+3])
		if i+3 == len(number) {
			break
		}
		buf.WriteByte(',')
	}

	buf.WriteByte(']')
	s := buf.String()
	println(s)
}
