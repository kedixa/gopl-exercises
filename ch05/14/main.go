package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
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

func main() {
	for i, course := range breadthFirst(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func breadthFirst(prereqs map[string][]string) []string {
	var worklist []string
	for key := range prereqs {
		worklist = append(worklist, key)
	}
	result := make([]string, 0, 50)
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				result = append(result, item)
				seen[item] = true
				worklist = append(worklist, prereqs[item]...)
			}
		}
	}
	return result
}
