// Crawl crawls web links starting with the command-line arguments.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yyamada12/go_exercises/ch08/ex06/links"
)

var maxDepth = flag.Int("depth", 3, "number of times crawling links")

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

type workItem struct {
	urlList []string
	depth   int
}

func main() {
	flag.Parse()
	worklist := make(chan workItem)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- workItem{flag.Args(), 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		work := <-worklist
		if work.depth > *maxDepth {
			return
		}
		for _, link := range work.urlList {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- workItem{crawl(link), work.depth + 1}
				}(link)
			}
		}
	}
}
