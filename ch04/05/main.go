package main

import (
	"fmt"
)

func unique(str []string) []string {
	idx := 1
	// for each string in str[1:], if not equal to
	// str[idx-1],  then assign to str[idx],
	// and increase idx
	for _, s := range str[1:] {
		if s != str[idx-1] {
			str[idx] = s
			idx++
		}
	}
	return str[:idx]
}

func main() {
	str := []string{"a", "a", "b", "c", "c", "a"}
	fmt.Println(unique(str))
}
