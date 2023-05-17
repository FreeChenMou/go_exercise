package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

//exercise4.2 根据情况输出散列字符
func main() {
	length := len(os.Args)
	if length == 3 {
		if strings.HasPrefix(os.Args[2], "384") {
			HandlerSHA384(os.Args[2])
		} else {
			HandlerSHA512(os.Args[2])
		}

	} else {
		sum256 := sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("%x", sum256)
	}
}

func HandlerSHA384(s string) {
	sum384 := sha512.Sum384([]byte(s))
	fmt.Printf("%x", sum384)
}

func HandlerSHA512(s string) {
	sum512 := sha512.Sum512([]byte(s))
	fmt.Printf("%x", sum512)
}

//exercise4.1 统计重复字符
func PopCount() {
	sum256 := sha256.Sum256([]byte("y"))
	record := make(map[byte]int)
	for i := 0; i < len(sum256); i++ {
		record[sum256[i]]++
	}
	var count int
	for _, b := range record {
		if b > 0 {
			count++
		}
	}
	println(count)
}
