// Mandelbrot emits a PNG image of the Mandelbrot fractal.
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
	dx, dy := 0.5/width*(xmax-xmin), 0.5/height*(ymax-ymin)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z1 := complex(x, y)
			z2 := complex(x+dx, y)
			z3 := complex(x, y+dy)
			z4 := complex(x+dx, y+dy)
			img.Set(px, py, averageColor(mandelbrot(z1), mandelbrot(z2), mandelbrot(z3), mandelbrot(z4)))
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func averageColor(colors ...color.Color) color.Color {
	var r, g, b, a uint32
	for _, c := range colors {
		cr, cg, cb, ca := c.RGBA()
		r += cr
		g += cg
		b += cb
		a += ca
	}
	l := uint32(len(colors))
	r /= l
	g /= l
	b /= l
	a /= l
	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{242 - n, 213 - 7*n, 15 + 3*n, 255}
		}
	}
	return color.Black
}
