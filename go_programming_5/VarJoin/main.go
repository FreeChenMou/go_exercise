package main

import (
	"fmt"
	"strings"
)

//exercise5.16 变长版stings.join
func main() {
	s := join(" ", "cnm")
	fmt.Println(s)

}

func join(sep string, str ...string) string {
	switch len(str) {
	case 0:
		return ""
	case 1:
		return str[0]
	}
	var builder strings.Builder

	n := len(sep) * len(str)
	for _, s := range str {
		n += len(s)
	}

	builder.Grow(n)

	for _, s := range str {
		builder.WriteString(s)
		builder.WriteString(sep)
	}

	return builder.String()
}
