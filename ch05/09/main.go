package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func expand(s string, f func(string) string) string {
	// find $... and expand to the result of f(...)
	var buf bytes.Buffer
	d := ""
	flag := false
	for _, r := range s {
		if flag {
			// find all the letters after $
			if unicode.IsLetter(r) {
				d += string(r)
				continue
			} else {
				buf.WriteString(f(d))
				d = ""
				flag = false
			}
		}
		// find a '$'
		if r == '$' {
			flag = true
			continue
		}
		buf.WriteRune(r)
	}
	if flag {
		buf.WriteString(f(d))
	}
	return buf.String()
}
func main() {
	fmt.Println(expand("$abc def $ddd ddd $xyz", strings.ToUpper))
}
