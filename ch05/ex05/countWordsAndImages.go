package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}
		fmt.Printf("%d words, %d images found in %s\n", words, images, url)
	}
}

// CountWordsAndImages fetchs url, parses html and counts words and images
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return visit(words, images, n)
}

func visit(words, images int, n *html.Node) (int, int) {

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		for _, line := range strings.Split(n.Data, "\n") {
			text := strings.TrimSpace(line)
			if text != "" {
				scanner := bufio.NewScanner(strings.NewReader(text))
				scanner.Split(bufio.ScanWords)
				for scanner.Scan() {
					words++
				}
				if err := scanner.Err(); err != nil {
					fmt.Fprintln(os.Stderr, "reading TextNode:", err)
				}
			}
		}

	}
	// skip if n is script or style tag
	if !(n.Type == html.ElementNode) || (n.Data != "script" && n.Data != "style") {
		if c := n.FirstChild; c != nil {
			words, images = visit(words, images, c)
		}
	}
	if c := n.NextSibling; c != nil {
		words, images = visit(words, images, c)
	}
	return words, images
}
