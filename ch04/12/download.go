package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type xkcdComics struct {
	Title      string
	Year       string
	Month      string
	Day        string
	Num        int
	Link       string
	News       string
	Transcript string
}

func download() error {
	allComics := make([]xkcdComics, 0, 2500)
	// use the first 100 for example
	for i := 1; i <= 100; i++ {
		fmt.Printf("\rdownloading %d", i)
		infoURL := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		resp, err := http.Get(infoURL)
		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusNotFound {
			// the last one
			resp.Body.Close()
			break
		} else if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return fmt.Errorf("download: get %s failed, %s", infoURL, resp.Status)
		}
		var result xkcdComics
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return err
		}
		resp.Body.Close()
		allComics = append(allComics, result)
	}
	fmt.Printf("\n")
	// to json
	data, err := json.MarshalIndent(allComics, "", "    ")
	if err != nil {
		return err
	}
	// write to data.out
	err = ioutil.WriteFile("data.out", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
