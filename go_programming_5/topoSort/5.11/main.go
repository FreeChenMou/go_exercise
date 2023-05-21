package main

import (
	"fmt"
)

var prereqs2 = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// 练习5.10： 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。

func main() {
	for i, course := range topoSort2(prereqs2) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) bool

	visitAll = func(items []string) bool {
		for _, item := range items {

			for _, v2 := range m[item] {
				for _, v3 := range m[v2] {
					if v3 == item {
						return true
					}
				}
			}
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
		return false
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	isCircle := visitAll(keys)
	if isCircle {
		fmt.Println("there is a circle graph!")
	}
	return order
}
