package main

import "strings"

//exercise5.9 函数替换返回值字符串
func main() {
	s := "$fooabcefg$foocba"

	println(expand(s, strstr))
}

func expand(s string, f func(string) string) string {
	str := f("foo")
	s = strings.ReplaceAll(s, "$foo", str)
	return s
}

func strstr(s string) string {
	return s
}
