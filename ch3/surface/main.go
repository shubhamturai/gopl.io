// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"image"
	"image/color"
	"image/gif"
	"os"
	"log"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var invalid bool = false

var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	//blackIndex = 1 // next color in palette
	greenIndex = 1 // next color in palette
	redIndex   = 2 //next color in palette
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	
	
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			anim := gif.GIF{LoopCount: 1}
			for i := 0; i < cells; i++ {
				for j := 0; j < cells; j++ {
					ax, ay := corner(i+1, j)
					bx, by := corner(i, j)
					cx, cy := corner(i, j+1)
					dx, dy := corner(i+1, j+1)
					if invalid {
						invalid = false
						break
					}

					rect := image.Rect(0, 0, cells, cells)
					img := image.NewPaletted(rect, palette)
					img.SetColorIndex(int(ax), int(ay), greenIndex)
					img.SetColorIndex(int(bx), int(by), greenIndex)
					img.SetColorIndex(int(cx), int(cy), greenIndex)
					img.SetColorIndex(int(dx), int(dy), greenIndex)
					anim.Image = append(anim.Image, img)
					fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				}
			}
			fmt.Println("</svg>")
			gif.EncodeAll(w, &anim)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	if x == 0.0 && y == 0.0 {
		invalid = true
	}
	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
