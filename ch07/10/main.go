package main

import (
	"fmt"
	"sort"
)

type palin []byte

func (p palin) Len() int           { return len(p) }
func (p palin) Less(i, j int) bool { return p[i] < p[j] }
func (p palin) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func isPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	b1 := []byte("i love you")
	b2 := []byte("abcdcba")
	fmt.Printf("%q is palindrome: %t\n", b1, isPalindrome(palin(b1)))
	fmt.Printf("%q is palindrome: %t\n", b2, isPalindrome(palin(b2)))
}
