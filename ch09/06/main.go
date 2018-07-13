package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
)

// split each row as a sub work
type item struct {
	Py    int
	Y     float64
	Color []color.Color
}

const numOfWorkers = 8
const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 10240, 10240
)

var done = make(chan int)
var result = make(chan item, 1000)
var inCh []chan item

func main() {
	for i := 0; i < numOfWorkers; i++ {
		inCh = append(inCh, make(chan item, 100))
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// save the result
	go func() {
		for i := 0; i < height; i++ {
			it := <-result
			for j := 0; j < width; j++ {
				img.Set(j, it.Py, it.Color[j])
			}
		}
		done <- 1
	}()

	// start workers
	for i := 0; i < numOfWorkers; i++ {
		chI := inCh[i]
		go func() {
			// deal each input item, send to result
			for it := range chI {
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, it.Y)
					it.Color = append(it.Color, mandelbrot(z))
				}
				result <- it
			}
		}()
	}
	// send items to each worker
	curChan := 0
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		inCh[curChan] <- item{py, y, nil}
		curChan++
		if curChan >= numOfWorkers {
			curChan = 0
		}
	}
	<-done
	// ignore the output
	//png.Encode(os.Stdout, img)
	fmt.Println("done")
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
