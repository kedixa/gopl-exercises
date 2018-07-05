package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	shaType := "256"
	if len(os.Args) != 1 {
		shaType = os.Args[1]
	}
	// read input
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// deal with different sha types
	switch shaType {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256(b))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384(b))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512(b))
	default:
		fmt.Printf("Unknown type: %s\n", shaType)
	}
}
