// convert converts temperature, length, weight
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		input := bufio.NewReader(os.Stdin)
		for {
			arg, err := input.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Fprintf(os.Stderr, "convert: %v\n", err)
				os.Exit(1)
			}
			arg = strings.TrimRight(arg, "\r\n")
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert: %v\n", err)
				os.Exit(1)
			}
			conv(t)
		}
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert: %v\n", err)
				os.Exit(1)
			}
			conv(t)
		}
	}
}

func conv(t float64) {
	fmt.Printf("%g°C = %g°F, %g°F = %g°C\n", t, float64(t*9/5+32), t, float64((t-32)*5/9))
	fmt.Printf("%gm = %gin, %gin = %gm\n", t, t*39.3700787, t, t*0.0254)
	fmt.Printf("%gkg = %glb, %glb = %gkg\n", t, t*2.2046226, t, t*0.4535924)
}
