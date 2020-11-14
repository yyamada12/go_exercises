// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

type test struct {
	url  string
	done bool
}

func incomingURLs() <-chan test {
	ch := make(chan test)
	go func() {
		for _, t := range []test{
			{"https://golang.org", false},
			{"https://godoc.org", true},
			{"https://play.golang.org", false},
			{"http://gopl.io", false},
			{"https://golang.org", false},
			{"https://godoc.org", false},
			{"https://play.golang.org", false},
			{"http://gopl.io", false},
		} {
			ch <- t
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done <-chan struct{}) (interface{}, error)
}

type result struct {
	value interface{}
	err   error
}

func Sequential(t *testing.T, m M) {
	for t := range incomingURLs() {
		done := make(chan struct{})
		res := make(chan result)
		start := time.Now()
		go func(url string) {
			value, err := m.Get(url, done)
			res <- result{value, err}
		}(t.url)
		if t.done {
			time.Sleep(150 * time.Microsecond)
			close(done)
		}
		r := <-res
		if r.err != nil {
			log.Print(r.err)
			continue
		}
		if r.value != nil {
			fmt.Printf("%s, %s, %d bytes\n",
				t.url, time.Since(start), len(r.value.([]byte)))
		} else {
			fmt.Printf("%s, %s, canceled\n",
				t.url, time.Since(start))
		}
	}
}

func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for t := range incomingURLs() {
		n.Add(1)
		go func(t test) {
			defer n.Done()
			done := make(chan struct{})
			res := make(chan result)
			start := time.Now()
			go func() {
				value, err := m.Get(t.url, done)
				res <- result{value, err}
			}()
			if t.done {
				time.Sleep(150 * time.Microsecond)
				close(done)
			}
			r := <-res
			if r.err != nil {
				log.Print(r.err)
				return
			}
			if r.value != nil {
				fmt.Printf("%s, %s, %d bytes\n",
					t.url, time.Since(start), len(r.value.([]byte)))
			} else {
				fmt.Printf("%s, %s, canceled\n",
					t.url, time.Since(start))
			}
		}(t)
	}
	n.Wait()
}
