package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	angle   = math.Pi / 6
	cells   = 100
	xyrange = 30.0
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	// web server
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		// default value
		width, height := 600, 320
		color := "FFFFFF"
		// read values from form
		if len(r.Form["width"]) > 0 {
			w, err := strconv.Atoi(r.Form["width"][0])
			if err != nil {
				log.Print(err)
			}
			width = w
		}
		if len(r.Form["height"]) > 0 {
			h, err := strconv.Atoi(r.Form["height"][0])
			if err != nil {
				log.Print(err)
			}
			height = h
		}
		if len(r.Form["color"]) > 0 {
			color = r.Form["color"][0]
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w, width, height, color)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(out io.Writer, width int, height int, color string) {
	// write the resule to out
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: #%s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j+1, width, height)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j, width, height int) (float64, float64) {
	zscale := float64(height) * 0.4
	xyscale := float64(width) / 2 / xyrange
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
