// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

//exercise4.8 统计字母、数字和其他unicode中的分类字符数
func main() {

	var letterCount, numberCount, otherCount, invalid int // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			letterCount++
			continue
		}
		if unicode.IsNumber(r) {
			numberCount++
			continue
		}

		otherCount++
	}

	fmt.Printf("letter:%d\nnumber:%d\nother:%d\ncount:%d",
		letterCount, numberCount, otherCount, invalid)

}
