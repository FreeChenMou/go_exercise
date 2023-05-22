package main

import "sort"

type s struct {
	s []string
}

func (s s) Len() int           { return len(s.s) }
func (s s) Less(i, j int) bool { return s.s[i] < s.s[j] }
func (s s) Swap(i, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }

func IsPalindrome(s sort.Interface) bool {
	var i, j = 0, s.Len() - 1

	for i < j {
		if !s.Less(i, j) && !s.Less(j, i) { // 索引i和j上的元素相等
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

//exercise7.10
func main() {

}
