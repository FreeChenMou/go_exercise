package main

import (
	"bytes"
	"fmt"
	"math"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {

	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
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

//exercise6.1 实现更多方法
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if (word & (1 << uint64(j))) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) remove(x int) {
	word, bit := x/64, x%64
	s.words[word] ^= uint64(1 << bit)
}

func (s *IntSet) Clear() {
	for i := 0; i < len(s.words); i++ {
		s.words[i] &= 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var t IntSet
	for i, word := range s.words {
		s.words[i] = word
	}
	return &t
}

func main() {
	set := &IntSet{}
	for i := 0; i < int(math.Pow(2, 10)); i += 64 {
		set.Add(i)
	}
	set.remove(64)
	println(set.Len())
}
