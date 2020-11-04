package main

import (
	"image"
	"image/png"
	"io/ioutil"
	"sync"
	"testing"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

var img *image.RGBA

func Benchmark_OnlyMainGoRoutine(b *testing.B) {
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

func Benchmark_WithWidthGoRoutine(b *testing.B) {

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

func Benchmark_WithHeightGoRoutine(b *testing.B) {

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

func Benchmark_WithWidthHeightGoRoutine(b *testing.B) {

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
