package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}

	// x: string before dot, y: string after dot, sign: sign
	var x, y, sign string
	if s[0] == '+' || s[0] == '-' {
		sign = s[:1]
		s = s[1:]
	}
	// find the place of dot, and split s to x and y
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		x, y = s[:dot], s[dot+1:]
	} else {
		x = s
	}
	nx, ny := len(x), len(y)
	// deal x
	var buf bytes.Buffer
	buf.WriteString(sign)
	i, j := 0, nx%3
	if j == 0 {
		j = 3
	}
	buf.WriteString(x[i:j])
	j += 3
	for j <= nx {
		buf.WriteByte(',')
		buf.WriteString(x[j-3 : j])
		j += 3
	}
	// deal y
	if dot >= 0 {
		buf.WriteString(".")
		j = 0
		for j < ny {
			if j+3 < ny {
				buf.WriteString(y[j : j+3])
				buf.WriteString(",")
			} else {
				buf.WriteString(y[j:ny])
			}
			j += 3
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234.5678"))
	fmt.Println(comma("-123.456"))
	fmt.Println(comma("+1234567890"))
	fmt.Println(comma("+1."))
	fmt.Println(comma("0.12345"))
}
