package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

//exercise6.2 多参数添加
func (s *IntSet) AddAll(arr ...int) {
	for _, x := range arr {
		word, bit := x/64, x%64
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	set := &IntSet{}
	set.AddAll(1, 2, 3, 4)
	s := set.String()
	println(s)
}
