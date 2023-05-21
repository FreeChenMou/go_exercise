// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

//!+table
//exercise5.10 prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

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

//!-table

//!+main
func main() {
	for key, value := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", key, value)
	}
}

func topoSort(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seen := make(map[string]bool)
	var visitAll func(items []string)
	index := 1
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[index] = item
				index++
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	visitAll(keys)
	return order
}
