// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

const (
	cells = 100         // number of grid cells
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			opts := handleQueryStrings(r.Form)
			w.Header().Set("Content-Type", "image/svg+xml")
			duplexFuncPlotter(w, opts)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	duplexFuncPlotter(os.Stdout, duplexFuncOpts{})
}

type duplexFuncOpts struct {
	width      int     // canvas width in pixels
	height     int     // canvas height in pixels
	xyrange    float64 // axis ranges (-xyrange..+xyrange)
	zscalecoef float64 // coefficient of pixels per z unit
	color      string  // color of image
}

func (opts *duplexFuncOpts) validateValues() {
	if opts.width <= 0 {
		opts.width = 600
	}
	if opts.height <= 0 {
		opts.height = 320
	}
	if opts.xyrange <= 0.0 {
		opts.xyrange = 30.0
	}
	if opts.zscalecoef <= 0.0 {
		opts.zscalecoef = 0.4
	}
	if !matchRegexp(`^[0-9a-f]{3}$`, opts.color) && !matchRegexp(`^[0-9a-f]{6}$`, opts.color) {
		opts.color = ""
	}
}

func matchRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

func handleQueryStrings(qs url.Values) duplexFuncOpts {
	opts := duplexFuncOpts{}
	for k, v := range qs {
		switch k {
		case "width":
			if w, err := strconv.Atoi(v[0]); err == nil {
				opts.width = w
			}
		case "height":
			if h, err := strconv.Atoi(v[0]); err == nil {
				opts.height = h
			}
		case "xyrange":
			if r, err := strconv.ParseFloat(v[0], 64); err == nil {
				opts.xyrange = r
			}
		case "zscalecoef":
			if z, err := strconv.ParseFloat(v[0], 64); err == nil {
				opts.zscalecoef = z
			}
		case "color":
			opts.color = v[0]
		default:
			log.Printf("invalid url parameter %q = %q", k, v)
		}
	}
	return opts
}

func duplexFuncPlotter(out io.Writer, opts duplexFuncOpts) {
	opts.validateValues()
	println("width: ", opts.width)
	println("height: ", opts.height)
	println("xyrange: ", opts.xyrange)
	println("zscalecoef: ", opts.zscalecoef)
	println("color: ", opts.color)

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", opts.width, opts.height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, err := corner(i+1, j, opts)
			if err != nil {
				continue
			}
			bx, by, bz, err := corner(i, j, opts)
			if err != nil {
				continue
			}
			cx, cy, cz, err := corner(i, j+1, opts)
			if err != nil {
				continue
			}
			dx, dy, dz, err := corner(i+1, j+1, opts)
			if err != nil {
				continue
			}
			z := int((az + bz + cz + dz) * 80)
			if z > 255 {
				z = 255
			} else if z < -255 {
				z = -255
			}
			color := opts.color
			if color == "" {
				if z > 0 {
					color = fmt.Sprintf("%02x0000", z)
				} else {
					color = fmt.Sprintf("0000%02x", -z)
				}
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"#%s\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i int, j int, opts duplexFuncOpts) (float64, float64, float64, error) {

	// Find point (x,y) at corner of cell (i,j).
	x := opts.xyrange * (float64(i)/cells - 0.5)
	y := opts.xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, errors.New("f(x,y) returned invalid value")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	xyscale := float64(opts.width) / 2 / opts.xyrange
	zscale := opts.zscalecoef * float64(opts.height)
	sx := float64(opts.width)/2 + (x-y)*cos30*xyscale
	sy := float64(opts.height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
