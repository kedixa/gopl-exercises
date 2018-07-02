package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// check weather url has http prefix
		if !strings.HasPrefix(url, "http://") {
			// if not, add http prefix
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// print the status of resp
		fmt.Println("Status: ", resp.Status)
	}
}
