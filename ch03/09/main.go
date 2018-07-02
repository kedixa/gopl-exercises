package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {

	const (
		width, height = 1024, 1024
	)
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		x, y := 2.0, 2.0
		if len(r.Form["x"]) > 0 {
			c, err := strconv.ParseFloat(r.Form["x"][0], 64)
			if err != nil {
				log.Print(err)
			}
			x = c
		}
		if len(r.Form["y"]) > 0 {
			c, err := strconv.ParseFloat(r.Form["y"][0], 64)
			if err != nil {
				log.Print(err)
			}
			y = c
		}
		xmin, ymin, xmax, ymax := -x, -y, x, y
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}
		png.Encode(w, img) // NOTE: ignoring errors
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
