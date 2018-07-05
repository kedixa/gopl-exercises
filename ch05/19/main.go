package main

import (
	"fmt"
)

// no return but return 1
func f() (x int) {
	defer func() {
		recover()
		x = 1
	}()
	panic("")
}

func main() {
	fmt.Println(f())
}
