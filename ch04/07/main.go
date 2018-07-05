package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(u []byte) []byte {
	lenu := len(u)
	for i := 0; i < lenu; {
		// for each rune
		_, sz := utf8.DecodeRune(u[i:])
		// reverse a rune
		switch sz {
		case 1:
			// do nothing
		case 2:
			u[i], u[i+1] = u[i+1], u[i]
		case 3:
			u[i], u[i+2] = u[i+2], u[i]
		case 4:
			u[i], u[i+1], u[i+2], u[i+3] = u[i+3], u[i+2], u[i+1], u[i]
		default:
			// no such rune
		}
		i += sz
	}
	// then, reverse u
	for i, j := 0, lenu-1; i < j; i, j = i+1, j-1 {
		u[i], u[j] = u[j], u[i]
	}
	return u
}

func main() {
	var u = []byte("hello, 世界")
	fmt.Printf("before reverse: %s\n", u)
	fmt.Printf("after reverse: %s\n", reverse(u))
}
