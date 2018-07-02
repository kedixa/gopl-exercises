package main

import (
	"fmt"
	"os"
)

func main() {
	for id, arg := range os.Args {
		fmt.Println(id, arg)
	}
}
