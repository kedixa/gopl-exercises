package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 4x supersampling
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		var row = make([]int, width)
		for px := 0; px < width; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(float64(px)+0.5)/width*(xmax-xmin) + xmin
			row[px] = mandelbrot(complex(x1, y)) + mandelbrot(complex(x2, y))
		}
		y = float64(float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(float64(px)+0.5)/width*(xmax-xmin) + xmin
			row[px] += mandelbrot(complex(x1, y)) + mandelbrot(complex(x2, y))
		}
		for px := 0; px < width; px++ {
			img.Set(px, py, color.Gray{uint8(float64(row[px]) / 4.0)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) int {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return 255 - contrast*int(n)
		}
	}
	return 0
}
