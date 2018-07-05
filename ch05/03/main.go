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
	visit(doc)
}

// visit, traverse all elements and output the text
func visit(n *html.Node) {
	// ignore script and style
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	// output the text
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	// visit all children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
