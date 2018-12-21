package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg" // register PNG decoder
	"image/png"
	"io"
	"os"
)

var imageType string

func main() {
	flag.StringVar(&imageType, "type", "jpeg", "-type jpeg")
	flag.Parse()
	if err := convert(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch imageType {
	case "jpeg":
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	}
	return fmt.Errorf("convert: Unknown image type")
}
