package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter count words
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		*c++
	}
	return len(p), nil
}

// line counter is the same as word counter
// ...

func main() {
	var c WordCounter
	c.Write([]byte("hello world\nI love you"))
	fmt.Println(c)
}
