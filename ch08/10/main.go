package main

import (
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// cancel cancel all the request if cancel is closed
var cancel = make(chan struct{})

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// input a byte, cancel all
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
		for range unseenLinks {
			// do nothing
		}
	}()
	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for {
				select {
				case link := <-unseenLinks:
					foundLinks := crawl(link)
					if foundLinks == nil {
						continue
					}
					go func() { worklist <- foundLinks }()
				case <-cancel:
					return
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
loop:
	for {
		select {
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		case <-cancel:
			close(unseenLinks)
			break loop
		}
	}
}
