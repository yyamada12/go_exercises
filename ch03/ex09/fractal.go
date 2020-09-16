package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			opts := handleQueryStrings(r.Form)
			fractal(w, opts)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	fractal(os.Stdout, fractalOpts{x: 0, y: 0, xyrange: 4.0})
}

type fractalOpts struct {
	x       float64 //
	y       float64 //
	xyrange float64 // range of x, y
}

func handleQueryStrings(qs url.Values) fractalOpts {
	opts := fractalOpts{}
	for k, v := range qs {
		switch k {
		case "x":
			if x, err := strconv.ParseFloat(v[0], 64); err == nil {
				opts.x = x
			} else {
				log.Printf("parse error with url parameter %q = %q", k, v)
			}
		case "y":
			if y, err := strconv.ParseFloat(v[0], 64); err == nil {
				opts.y = y
			} else {
				log.Printf("parse error with url parameter %q = %q", k, v)
			}
		case "xyrange":
			if s, err := strconv.ParseFloat(v[0], 64); err == nil {
				if s <= 0.0 {
					log.Printf("xyrange must be positive value, invalid %q = %q", k, v)
				} else {
					opts.xyrange = s
				}
			} else {
				log.Printf("parse error with url parameter %q = %q", k, v)
			}
		default:
			log.Printf("invalid url parameter %q = %q", k, v)
		}

	}
	return opts
}

func fractal(out io.Writer, opts fractalOpts) {

	const width, height = 1024, 1024

	xmin := opts.x - opts.xyrange/2.0
	ymin := opts.y - opts.xyrange/2.0
	xmax := opts.x + opts.xyrange/2.0
	ymax := opts.y + opts.xyrange/2.0

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

	png.Encode(out, img) // NOTE: ignoring errors
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
