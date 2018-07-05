package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func elementsByTagName(doc *html.Node, name ...string) []*html.Node {
	// put the names into a map, so we can easily judge whether a name exists
	mp := make(map[string]bool)
	for _, n := range name {
		mp[n] = true
	}
	// recursive visit, store to result
	result := visit(doc, nil, mp)
	return result
}

func visit(n *html.Node, result []*html.Node, name map[string]bool) []*html.Node {
	if n.Type == html.ElementNode && name[n.Data] {
		result = append(result, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = visit(c, result, name)
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s https://your.url", os.Args[0])
		return
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	imgs := elementsByTagName(doc, "img")
	fmt.Printf("%d images\n", len(imgs))
	hds := elementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Printf("%d headings\n", len(hds))
}
