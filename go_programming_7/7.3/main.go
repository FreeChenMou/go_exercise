package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

//exercise7.3 输出树节点
func (t *tree) String() string {
	str := ""
	if t == nil {
		return str
	} else {
		str += t.left.String()
		str = fmt.Sprintf("%d", t.value)
		str += t.right.String()
	}
	return str
}

func main() {
	t := &tree{value: 10, left: nil, right: nil}
	println(t.String())
}
