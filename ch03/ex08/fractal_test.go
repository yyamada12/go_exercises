package main

import (
	"math/big"
	"testing"
)

const (
	// for big.Rat
	xminNum, yminNum, xmaxNum, ymaxNum = 97, 0, 100, 3
	denom                              = 100

	xmin, ymin, xmax, ymax = 0.97, 0, 1, 0.03
	width, height          = 30, 30
)

func BenchmarkNewtonWithComplex64(b *testing.B) {
	iterations = 4
	for i := 0; i < b.N; i++ {
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				newtonWithComplex64(x, y)
			}
		}
	}
}

func BenchmarkNewtonWithComplex128(b *testing.B) {
	iterations = 4
	for i := 0; i < b.N; i++ {
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				newtonWithComplex128(x, y)
			}
		}
	}
}

func BenchmarkNewtonWithComplexFloat(b *testing.B) {
	iterations = 4
	for i := 0; i < b.N; i++ {
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				newtonWithComplexFloat(x, y)
			}
		}
	}
}

func BenchmarkNewtonWithComplexRat(b *testing.B) {
	iterations = 4
	for i := 0; i < b.N; i++ {
		for py := 0; py < height; py++ {
			y := big.NewRat(int64(py)*(ymaxNum-yminNum)+yminNum*height, height*denom)
			for px := 0; px < width; px++ {
				x := big.NewRat(int64(px)*(xmaxNum-xminNum)+xminNum*width, width*denom)
				newtonWithComplexRat(x, y)
			}
		}
	}
}
