package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v", arg, err)
			continue
		}
		if hasDup(f) {
			fmt.Println(arg)
		}
		f.Close()
	}
}

func hasDup(f *os.File) bool {
	input := bufio.NewScanner(f)
	counts := make(map[string]int)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if counts[line] > 1 {
			return true
		}
	}
	return false
}
