package main

import (
	"fmt"
)

// refer to http://www.cplusplus.com/reference/algorithm/rotate/
// I think I haven’t understood yet
func rotate(s []int, n int) {
	lens := len(s)
	if lens == 0 {
		return
	}
	n = n % lens
	first, next := 0, n
	for first != next {
		s[first], s[next] = s[next], s[first]
		first, next = first+1, next+1
		if next == lens {
			next = n
		} else if first == n {
			n = next
		}
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 3)
	fmt.Println(a)
}
