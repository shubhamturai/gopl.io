// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
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
	if len(os.Args[:1]) > 0{
		for _, imgStr := range os.Args[1:]{
			f, err := os.Create(imgStr)
			if err != nil {
				fmt.Println("could not create an empty image")
			}
			defer f.Close()
			img := image.NewRGBA(image.Rect(0, 0, width, height))
			var mandelColor color.Color
			for py := 0; py < height; py += 2 { //originally it was py++
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px += 2 { //originally it was px++
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					//img.Set(px, py, mandelbrot(z))
					//mandelColor = mandelbrot(z)
					mandelColor = newton(z)
					img.Set(px, py, mandelColor)
					img.Set(px+1, py, mandelColor)
					img.Set(px, py+1, mandelColor)
					img.Set(px+1, py+1, mandelColor)
				}
			}
			//png.Encode(os.Stdout, img) // NOTE: ignoring errors //gives output to the console
			
			// Encode to `PNG` with `DefaultCompression` level
			// then save to file
			err = png.Encode(f, img)
			if err != nil {
				fmt.Println("could not encode the created image to png format")
			}
		}
		
	} else{
		fmt.Print("please give the names of the images as an argument while executing")
	}
	
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			return color.RGBA{255 - contrast*n, 0x00, 0x00, 0xff}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			//return color.Gray{255 - contrast*i}
			return color.RGBA{255 - contrast*i, 0x00, 0x00, 0xff}
		}
	}
	return color.Black
}
