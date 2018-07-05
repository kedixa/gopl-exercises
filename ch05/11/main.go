package main

import (
	"fmt"
	"os"
	"sort"
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
	"linear algebra":        {"calculus"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	// seen[item] == 1 means we see item but not satisfied
	seen := make(map[string]int)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if seen[item] == 0 {
				seen[item] = 1
				visitAll(m[item])
				seen[item] = 2
				order = append(order, item)
			} else if seen[item] == 1 {
				// if the current dependency is seen but not satisfied, there is a circle
				fmt.Println("Circle: ", item)
				os.Exit(1)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
