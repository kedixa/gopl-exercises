package main

import (
	"fmt"
)

func join(sep string, vals ...string) (result string) {
	if len(vals) == 0 {
		return
	}
	result += vals[0]
	for _, v := range vals[1:] {
		result += sep
		result += v
	}
	return
}

func main() {
	fmt.Println(join(" ", "I", "LOVE", "YOU"))
}
