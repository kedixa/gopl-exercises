package main

import (
	"fmt"
)

func max(x int, vals ...int) int {
	result := x
	for _, v := range vals {
		if result < v {
			result = v
		}
	}
	return result
}

func min(x int, vals ...int) int {
	result := x
	for _, v := range vals {
		if result > v {
			result = v
		}
	}
	return result
}

func main() {
	fmt.Printf("max(5, 1, 6, 4): %d\n", max(5, 1, 6, 4))
	fmt.Printf("min(5, 1, 6, 4): %d\n", min(5, 1, 6, 4))
}
