package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var cancel = make(chan struct{})

func fetch(url string) struct{} {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	// if cancel if closed, req will be Canceled
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
	return struct{}{}
}

func main() {
	// fetch each url
	done := make(chan struct{}, len(os.Args))
	for _, url := range os.Args[1:] {
		go func(url string) { done <- fetch(url) }(url)
	}
	// until one of them is done
	<-done
	close(cancel)
}
