package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// get an apikey from http://www.omdbapi.com/
const (
	apikey = ""
)

type movie struct {
	Title  string
	Poster string
}

func main() {
	// join search words
	q := url.QueryEscape(strings.Join(os.Args[1:], " "))
	URL := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", apikey, q)
	// get the search result
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Printf("search failed: %s\n", resp.Status)
		os.Exit(1)
	}
	var result movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Println(err)
		os.Exit(1)
	}
	resp.Body.Close()
	// print the url of poster
	fmt.Printf("Title: %s\nPoster URL: %s\n", result.Title, result.Poster)
}

// for example: ./main avengers infinity war
