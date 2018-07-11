package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Usage: ./main [name or attrs] ...
 * For example: ./main html id:first class:class
 */
func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", join(stack, " "), tok)
			}
		}
	}
}

// join join the names of xml.StartElement
func join(s []xml.StartElement, sep string) string {
	buf := bytes.NewBuffer(nil)
	if len(s) == 0 {
		return buf.String()
	}
	buf.WriteString(s[0].Name.Local)
	for _, t := range s[1:] {
		buf.WriteString(sep)
		buf.WriteString(t.Name.Local)
	}
	return buf.String()
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		name, value := "name", y[0]
		if strings.HasPrefix(y[0], "id:") {
			name = "id"
			value = y[0][3:]
		} else if strings.HasPrefix(y[0], "class:") {
			name = "class"
			value = y[0][6:]
		}
		// ... and some other attrs

		// check for name and attrs
		if name == "name" {
			if x[0].Name.Local == value {
				y = y[1:]
			}
		} else {
			for _, v := range x[0].Attr {
				if v.Name.Local == name {
					if v.Value == value {
						y = y[1:]
					}
					break
				}
			}
		}

		x = x[1:]
	}
	return false
}
