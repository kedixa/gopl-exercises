package main

import (
	"flag"
	"fmt"
	"log"
)

type item struct {
	depth int
	url   string
}

// max depth
var depth int

func crawl(it item) []item {
	if it.depth > depth {
		return nil
	}
	url := it.url
	fmt.Println(it.depth, url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	itList := make([]item, len(list))
	for i := 0; i < len(list); i++ {
		itList[i] = item{it.depth + 1, list[i]}
	}
	return itList
}

func main() {
	// parse flags
	flag.IntVar(&depth, "depth", 3, "-depth=3")
	flag.Parse()
	worklist := make(chan []item)  // lists of URLs, may have duplicates
	unseenLinks := make(chan item) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		var itList []item
		// other args are urls
		for _, i := range flag.Args() {
			itList = append(itList, item{0, i})
		}
		worklist <- itList
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}
