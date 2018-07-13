package main

import (
	"fmt"
)

const m = 1000000

func main() {
	ch := make([]chan int, m)
	for i := 0; i < m; i++ {
		ch[i] = make(chan int)
	}
	for i := 0; i < m-1; i++ {
		go func(in <-chan int, out chan<- int) {
			out <- <-in
		}(ch[i], ch[i+1])
	}
	ch[0] <- 123
	fmt.Println(<-ch[m-1])
}
