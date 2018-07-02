package main

import (
	"fmt"
)

func check(x, y string) bool {
	var ax, ay [256]int
	for c := range x {
		ax[int(c)]++
	}
	for c := range y {
		ay[int(c)]++
	}
	for i := 0; i < 256; i++ {
		if ax[i] != ay[i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(check("12345", "13524"))
	fmt.Println(check("12345", "135246"))
}
