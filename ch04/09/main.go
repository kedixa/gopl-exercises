package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		// read a word and count it
		w := input.Text()
		freq[w]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	// print each word and it's freq
	fmt.Println("word\tcount")
	for w, c := range freq {
		fmt.Printf("%s\t%d\n", w, c)
	}
}
