package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}

	// buffer to store string
	var buf bytes.Buffer
	// before the first comma
	i, j := 0, n%3
	if j == 0 {
		j = 3
	}
	buf.WriteString(s[i:j])
	j += 3
	// 3 digits and a comma
	for j <= n {
		buf.WriteByte(',')
		buf.WriteString(s[j-3 : j])
		j += 3
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}
