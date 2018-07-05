package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func mergeSpace(u []byte) []byte {
	lenu := len(u)
	if lenu == 0 {
		return u
	}
	preIsSpace := false // record whether the previous character is a space
	idx := 0            // dst index
	for i := 0; i < lenu; {
		// for each rune
		r, sz := utf8.DecodeRune(u[i:])
		// is r is space
		if unicode.IsSpace(r) {
			if !preIsSpace {
				u[idx] = ' '
				idx++
				preIsSpace = true
			}
		} else {
			// else copy the rune
			copy(u[idx:], string(r))
			idx += sz
			preIsSpace = false
		}
		i += sz
	}
	return u[:idx]
}
func main() {
	var r = []byte("  		你好 	 世  界   ")
	fmt.Printf("%s\n", string(mergeSpace(r)))
}
