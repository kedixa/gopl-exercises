package main

import (
	"bytes"
	"fmt"
	"io"
)

// CWriter counting writer
type CWriter struct {
	w *io.Writer
	c *int64
}

func (cw CWriter) Write(p []byte) (int, error) {
	*(cw.c) += int64(len(p))
	return (*(cw.w)).Write(p)
}

// CountingWriter construct a writer that can count the number of bytes
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c int64
	cw := CWriter{&w, &c}
	return cw, &c
}

func main() {
	w := bytes.NewBuffer(nil)
	cw, p := CountingWriter(w)
	cw.Write([]byte("I Love You\nGo Programming Language"))
	fmt.Printf("buf: %s\nCount: %d\n", w.String(), *p)
}
