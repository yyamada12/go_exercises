package main

import (
	"image"
	"image/png"
	"io/ioutil"
	"strconv"
	"sync"
	"testing"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

var img *image.RGBA

func Benchmark_OnlyMainGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img = image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}
		png.Encode(ioutil.Discard, img)
	}
}

func Benchmark_WithWidthGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img = image.NewRGBA(image.Rect(0, 0, width, height))
		var wg sync.WaitGroup
		for py := 0; py < height; py++ {
			wg.Add(1)
			go func(py int) {
				defer wg.Done()
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}(py)
		}
		wg.Wait()
		png.Encode(ioutil.Discard, img)
	}
}

func Benchmark_WithHeightGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img = image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			var wg sync.WaitGroup
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.

				wg.Add(1)
				go func(px, py int, z complex128) {
					defer wg.Done()
					img.Set(px, py, mandelbrot(z))
				}(px, py, z)
			}
			wg.Wait()
		}
		png.Encode(ioutil.Discard, img)
	}
}

func Benchmark_WithWidthHeightGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img = image.NewRGBA(image.Rect(0, 0, width, height))
		var wg sync.WaitGroup
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.

				wg.Add(1)
				go func(px, py int, z complex128) {
					defer wg.Done()
					img.Set(px, py, mandelbrot(z))
				}(px, py, z)
			}
		}
		wg.Wait()
		png.Encode(ioutil.Discard, img)
	}
}

func withCountingSemaphore(b *testing.B, n int) {
	tokens := make(chan struct{}, n)
	for i := 0; i < b.N; i++ {
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		var wg sync.WaitGroup
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.

				wg.Add(1)
				go func(px, py int, z complex128) {
					defer wg.Done()
					tokens <- struct{}{} // acquire a token
					img.Set(px, py, mandelbrot(z))
					<-tokens // release the token
				}(px, py, z)
			}
		}
		wg.Wait()
		png.Encode(ioutil.Discard, img)
	}
}

func Benchmark_withCountingSemaphore(b *testing.B) {
	for _, n := range []int{2, 5, 10, 20, 40, 80, 160, 320} {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			withCountingSemaphore(b, n)
		})
	}
}
