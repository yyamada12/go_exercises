// countHTMLElements prints counts of HTML element in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countHTMLElements: %v\n", err)
		os.Exit(1)
	}
	for element, count := range visit(map[string]int{}, doc) {
		fmt.Printf("%s\t%d\n", element, count)
	}
}

// visit counts up element found in n and returns the result.
func visit(elementCounts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elementCounts[n.Data]++
	}
	if c := n.FirstChild; c != nil {
		elementCounts = visit(elementCounts, c)
	}
	if c := n.NextSibling; c != nil {
		elementCounts = visit(elementCounts, c)
	}
	return elementCounts
}
