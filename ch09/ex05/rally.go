package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var n int64
	ping, pong := make(chan struct{}), make(chan struct{})
	go func() {
		for {
			<-ping
			atomic.AddInt64(&n, 1)
			pong <- struct{}{}
		}
	}()
	go func() {
		for {
			<-pong
			atomic.AddInt64(&n, 1)
			ping <- struct{}{}
		}
	}()
	ping <- struct{}{}
	time.Sleep(1 * time.Second)
	fmt.Printf("%d rallies during 1 second\n", atomic.LoadInt64(&n))
}
