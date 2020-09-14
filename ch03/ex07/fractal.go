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
			img.Set(px, py, averageColor(newton(z1), newton(z2), newton(z3), newton(z4)))
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

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{242 - i, 213 - 7*i, 15 + 3*i, 255}
		}
	}
	return color.Black
}
