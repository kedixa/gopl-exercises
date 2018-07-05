package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
func main() {
	// check whether data.out exists, or download
	if exists, err := fileExists("./data.out"); exists == false {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = download()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	// an example to search by year and/or month
	var year, month string
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--year" {
			i++
			year = os.Args[i]
		} else if os.Args[i] == "--month" {
			i++
			month = os.Args[i]
		}
	}
	// read the data
	data, err := ioutil.ReadFile("data.out")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var allComics []xkcdComics
	if err = json.Unmarshal(data, &allComics); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// show all comics that match the search criteria
	fmt.Println("------------------------------")
	for _, item := range allComics {
		if year != "" && item.Year != year {
			continue
		}
		if month != "" && item.Month != month {
			continue
		}
		fmt.Printf("URL: https://xkcd.com/%d/\nTranscript: %s\n", item.Num, item.Transcript)
		fmt.Println("------------------------------")
	}
}
