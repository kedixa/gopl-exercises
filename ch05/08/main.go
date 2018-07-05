package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

const id = "toc"

func main() {
	for _, url := range os.Args[1:] {
		node, err := outline(url)
		if err == nil {
			fmt.Printf("%v\n", node)
		}
	}
}

func outline(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return elementByID(doc, id), nil
}

func elementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, startElement, endElement, id)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, id string) bool, id string) *html.Node {
	if pre != nil {
		// if find the id in n
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post, id)
		if node != nil {
			return node
		}
	}

	if post != nil {
		post(n, id)
	}
	return nil
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	return false
}
