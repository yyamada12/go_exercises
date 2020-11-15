package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_pipeline(t *testing.T) {
	for _, n := range []int{1000000, 2000000, 4000000, 8000000} {
		measure(n)
	}
}

func measure(n int) {
	in, out := pipeline(n)
	start := time.Now()
	in <- struct{}{}
	<-out
	fmt.Printf("send struct{}{} through %d goroutines: %s\n", n, time.Since(start))
}
