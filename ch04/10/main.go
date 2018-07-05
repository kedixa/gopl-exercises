package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		// show issues one year ago
		if time.Since(item.CreatedAt).Hours() > 365*24 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
