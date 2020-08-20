package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		fetch(url)
	}
}

func fetch(url string) {

	resp, err := http.Get(complementURL(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(out, resp.Body); err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}

func complementURL(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "http://" + url
}
