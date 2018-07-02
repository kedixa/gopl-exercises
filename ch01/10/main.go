package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Using ./main prefix url...")
		os.Exit(1)
	}
	start := time.Now()
	ch := make(chan string)
	fprefix := os.Args[1] // add a fprefix to distinguish each run
	for id, url := range os.Args[2:] {
		go fetch(url, fmt.Sprint(fprefix, id, ".txt"), ch) // start a goroutine
	}
	for range os.Args[2:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// add a file name parameter
func fetch(url string, fname string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		ch <- fmt.Sprintf("Open %s failed.\n", fname)
		os.Exit(1)
		return
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	f.Close() // close the file
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
