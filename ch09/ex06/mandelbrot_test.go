package main

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"runtime"
	"strconv"
	"testing"
)

var img *image.RGBA

func Benchmark_createMandelbrotImage(b *testing.B) {
	cpus := runtime.NumCPU()
	fmt.Printf("cpu: %d\n", cpus)
	for i := 1; i <= cpus; i++ {
		b.Run("GOMAXPROCS:"+strconv.Itoa(i), func(b *testing.B) {
			runtime.GOMAXPROCS(i)
			for j := 0; j < b.N; j++ {
				img := createMandelbrotImage()
				png.Encode(ioutil.Discard, img)
			}
		})
	}
}
