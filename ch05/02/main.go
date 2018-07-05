package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	// save the result to a map
	dic := make(map[string]int)
	visit(dic, doc)
	for k, v := range dic {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

// visit, traverse all elements and count the number
func visit(dic map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		dic[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(dic, c)
	}
}
