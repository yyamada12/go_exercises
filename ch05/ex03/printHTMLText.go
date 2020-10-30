// printHTMLText prints text in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to texts each text found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		for _, line := range strings.Split(n.Data, "\n") {
			text := strings.TrimSpace(line)
			if text != "" {
				links = append(links, text)
			}
		}
	}
	// skip if n is script or style tag
	if !(n.Type == html.ElementNode) || (n.Data != "script" && n.Data != "style") {
		if c := n.FirstChild; c != nil {
			links = visit(links, c)
		}
	}
	if c := n.NextSibling; c != nil {
		links = visit(links, c)
	}
	return links
}
