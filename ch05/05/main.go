package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndImages ...
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {
	visit(n, &words, &images)
	return
}
func visit(n *html.Node, words, images *int) {
	// count images
	if n.Type == html.ElementNode && n.Data == "img" {
		*images++
	}
	// count words
	if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			*words++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, images)
	}
}

func main() {
	url := "https://www.golang.org"
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("words: %d, images: %d\n", words, images)
}
