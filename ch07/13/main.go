package main

import (
	"fmt"

	"./eval"
)

// the eval.Formatn did it
func main() {
	expr, _ := eval.Parse("1*2+3/pi+f(a,b)")
	fmt.Print(eval.Format(expr))
}
