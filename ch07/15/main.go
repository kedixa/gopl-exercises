package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./eval"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please input an expression: ")
		if !input.Scan() {
			fmt.Println("Good Bye!")
			return
		}
		// read expr
		expr := input.Text()
		e, err := eval.Parse(expr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// variables
		m := make(map[eval.Var]float64)
		m2 := make(map[eval.Var]bool)
		eval.Variables(e, m)
		for k := range m {
			for {
				// input variables
				fmt.Printf("Please input the value of %s: ", string(k))
				if !input.Scan() {
					// end of file
					return
				}
				v, err := strconv.ParseFloat(input.Text(), 64)
				if err != nil {
					fmt.Println(err)
					continue
				}
				m[k] = v
				m2[k] = true
				break
			}
		}
		err = e.Check(m2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := e.Eval(m)
		fmt.Printf("The result of %s is: %f\n", expr, result)
	}
}
