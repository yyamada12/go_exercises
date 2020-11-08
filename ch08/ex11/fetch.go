// Fetch prints the content found at each specified URL.
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// mirroredQuery(os.Args[1:])
	fmt.Printf("%s", mirroredQuery(os.Args[1:]))
}

type response struct {
	body []byte
	err  error
}

func mirroredQuery(urls []string) []byte {
	cancel := make(chan struct{})
	responses := make(chan response, len(urls))

	for _, url := range urls {
		go func(url string) {
			query(url, responses, cancel)
		}(url)
	}
	for range urls {
		res := <-responses
		if res.err != nil {
			log.Print(res.err)
		} else {
			close(cancel)
			return res.body
		}
	}
	return nil
}

func query(url string, responses chan<- response, cancel <-chan struct{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		responses <- response{nil, err}
		return
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	req = req.WithContext(ctx)

	go func() {
		<-cancel
		cancelFunc()
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		responses <- response{nil, err}
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	responses <- response{b, err}
}
