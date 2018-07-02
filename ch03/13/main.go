package main

import "fmt"

// KB to YB, I don't think itoa can be used in this example
const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Printf("KB: %d, MB: %d ...", KB, MB)
}
