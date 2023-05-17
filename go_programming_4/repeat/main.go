package main

import "fmt"

//exercise4.5 去除相同字符串
func main() {
	str := []string{"hahah", "abc", "abc"}
	strings := handle(str)

	for _, s := range strings {
		fmt.Printf("%s ", s)
	}
}

func handle(str []string) []string {
	for i, j := 0, 1; j < len(str); i, j = i+1, j+1 {
		if str[i] == str[j] {
			copy(str[i:], str[j:])
			str = str[:len(str)-1]
			i--
			j--
		}
	}
	return str
}
