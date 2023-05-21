package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

//exercise6.3 交集
func (s1 *IntSet) Intersection(s2 *IntSet) *IntSet {
	len1, len2 := len(s1.words), len(s2.words)

	len := len2
	if len1 < len2 {
		len = len1
	}

	t := &IntSet{}

	for i := 0; i < len; i++ {
		temp := s1.words[i] & s2.words[i]
		t.words = append(t.words, temp)
	}

	return t
}

//exercise6.3 差集
func (s1 *IntSet) DifferenceSet(s2 *IntSet) *IntSet {
	t := &IntSet{}

	for i, word := range s1.words {
		if i < len(s2.words) {
			t.words = append(t.words, word^s2.words[i])
		} else {
			t.words = append(t.words, word)
		}
	}

	return t
}

func (s1 *IntSet) SymmetricDifference(s2 *IntSet) *IntSet {
	set1 := s1.DifferenceSet(s2)
	set2 := s2.DifferenceSet(s1)
	set1.UnionWith(set2)
	return set1
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

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if (word & (1 << uint(j))) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) AddAll(arr ...int) {
	for _, x := range arr {
		word, bit := x/64, x%64
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

func main() {
	s1, s2 := &IntSet{}, &IntSet{}
	s1.AddAll(1, 2, 3, 4, 65)
	s2.AddAll(1, 2, 3, 4, 54)
	symmetricDifference := s1.SymmetricDifference(s2)
	println(symmetricDifference.String()) //{54 65}
}
