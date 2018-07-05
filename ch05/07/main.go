package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// test code ignored, run using ./main https://your.url
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	// Element
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		// attributes
		for _, a := range n.Attr {
			fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
		}
		// whether there are child nodes
		if n.FirstChild != nil {
			fmt.Printf(">\n")
		}
		depth++
	} else if n.Type == html.TextNode {
		// scan lines of the text
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanLines)
		// for each line, output with indent
		for input.Scan() {
			text := input.Text()
			text = strings.Trim(text, " \t\n")
			if len(text) > 0 {
				fmt.Printf("%*s%s\n", depth*2, "", text)
			}
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
			fmt.Printf("/>\n")
		} else {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	} else if n.Type == html.TextNode {
	}
}
