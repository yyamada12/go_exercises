package main

import "fmt"

func main() {
	in, out := pipeline(10)
	in <- 4
	fmt.Println(<-out)
}

func pipeline(n int) (chan<- interface{}, <-chan interface{}) {
	var ch <-chan interface{}
	in := make(chan interface{})
	ch = in
	for i := 0; i < n; i++ {
		ch = job(ch)
	}
	return in, ch
}

func job(prev <-chan interface{}) <-chan interface{} {
	next := make(chan interface{})
	go func() {
		val := <-prev
		next <- val
	}()
	return next
}
