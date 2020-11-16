package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	ping, pong := make(chan struct{}), make(chan struct{})
	go func() {
		for {
			<-ping
			n++
			pong <- struct{}{}
		}
	}()
	go func() {
		for {
			<-pong
			ping <- struct{}{}
		}
	}()
	ping <- struct{}{}
	time.Sleep(1 * time.Second)
	fmt.Printf("%d rallies during 1 second\n", n)
}
