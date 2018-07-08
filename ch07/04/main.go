package main

import (
	"fmt"
	"io"
)

// StrReader string reader, implement several functions of strings.Reader
type StrReader struct {
	data []byte
	curr int
}

// Len the number of remaining characters
func (r *StrReader) Len() int {
	return len(r.data) - r.curr
}

// ReadByte read one byte
func (r *StrReader) ReadByte() (b byte, err error) {
	if r.curr >= len(r.data) {
		err = io.EOF
	} else {
		b = r.data[r.curr]
		r.curr++
	}
	return
}

// Read read to []byte
func (r *StrReader) Read(b []byte) (n int, err error) {
	if r.curr >= len(r.data) {
		err = io.EOF
		return
	}
	// read bytes to b, at most len(b) bytes
	for i := 0; i < len(b); i++ {
		if r.curr < len(r.data) {
			b[i] = r.data[r.curr]
			r.curr++
			n++
		} else {
			break
		}
	}
	return
}
func newReader(s string) *StrReader {
	r := StrReader{[]byte(s), 0}
	return &r
}

func main() {
	r := newReader("I love you")
	fmt.Printf("Len: %d\n", r.Len())
	b, err := r.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("byte: %c\n", b)

	bs := make([]byte, 9)
	n, err := r.Read(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n, string(bs[:n]))
	b, err = r.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
}
