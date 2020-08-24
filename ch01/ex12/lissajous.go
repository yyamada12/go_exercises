package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			opts := handleQueryStrings(r.Form)
			lissajousWithOpts(w, opts)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

type lissajousOpts struct {
	cycles  int     // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
}

func (opts *lissajousOpts) validateValues() {
	if opts.cycles <= 0 {
		opts.cycles = 5
	}
	if opts.res <= 0.0 {
		opts.res = 0.001
	}
	if opts.size <= 0 {
		opts.size = 100
	}
	if opts.nframes <= 0 {
		opts.nframes = 64
	}
	if opts.delay <= 0 {
		opts.delay = 8
	}
}

func handleQueryStrings(qs url.Values) lissajousOpts {
	opts := lissajousOpts{}
	for k, v := range qs {
		switch k {
		case "cycles":
			if c, err := strconv.Atoi(v[0]); err == nil {
				opts.cycles = c
			}
		case "res":
			if r, err := strconv.ParseFloat(v[0], 64); err == nil {
				opts.res = r
			}
		case "size":
			if s, err := strconv.Atoi(v[0]); err == nil {
				opts.size = s
			}
		case "nframes":
			if nf, err := strconv.Atoi(v[0]); err == nil {
				opts.nframes = nf
			}
		case "delay":
			if d, err := strconv.Atoi(v[0]); err == nil {
				opts.delay = d
			}
		default:
			log.Printf("invalid url parameter %q = %q", k, v)
		}

	}
	return opts
}

func lissajousWithOpts(out io.Writer, opts lissajousOpts) {
	opts.validateValues()

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: opts.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < opts.nframes; i++ {
		rect := image.Rect(0, 0, 2*opts.size+1, 2*opts.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(opts.cycles)*2*math.Pi; t += opts.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				opts.size+int(x*float64(opts.size)+0.5),
				opts.size+int(y*float64(opts.size)+0.5),
				uint8(int(math.Floor(t))%7+1),
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, opts.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors

}

func lissajous(out io.Writer) {
	lissajousWithOpts(out, lissajousOpts{})
}
