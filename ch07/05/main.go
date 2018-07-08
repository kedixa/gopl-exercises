package main

import (
	"bytes"
	"fmt"
	"io"
)

type limitedReader struct {
	r io.Reader
	n int
}

// Read read bytes to b, at most len(b)
func (lr *limitedReader) Read(b []byte) (x int, err error) {
	// EOF
	if lr.n <= 0 {
		err = io.EOF
		return
	}
	// remain no less than len(b) bytes
	if len(b) <= lr.n {
		x, err = lr.r.Read(b)
		lr.n -= x
		return
	}
	// remain less than len(b) bytes
	x, err = lr.r.Read(b[:lr.n])
	lr.n -= x
	return
}

func limitReader(r io.Reader, n int) io.Reader {
	return &limitedReader{r, n}
}
func main() {
	// read at most 6 characters
	r := limitReader(bytes.NewReader([]byte("I love you")), 6)
	b := make([]byte, 7, 7)
	n, err := r.Read(b)
	fmt.Println(n, err)
	n, err = r.Read(b)
	fmt.Println(n, err)
}
