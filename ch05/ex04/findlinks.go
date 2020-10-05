// Findlinks prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for kind, links := range visit(map[string][]string{}, doc) {
		fmt.Println("<" + kind + ">")
		for _, link := range links {
			fmt.Println(link)
		}
		fmt.Println()
	}
}

// visit appends to links each link found in n and returns the result.
func visit(linkMap map[string][]string, n *html.Node) map[string][]string {
	if n.Type == html.ElementNode {

		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links, ok := linkMap["hyperlink"]
					if !ok {
						links = []string{}
					}
					linkMap["hyperlink"] = append(links, a.Val)
				}
			}
		case "link":
			var isCSS bool
			var href string
			for _, a := range n.Attr {
				if a.Key == "rel" {
					isCSS = a.Val == "stylesheet"
				} else if a.Key == "href" {
					href = a.Val
				}
			}
			if isCSS && href != "" {
				links, ok := linkMap["stylesheet"]
				if !ok {
					links = []string{}
				}
				linkMap["stylesheet"] = append(links, href)
			}
		case "script":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links, ok := linkMap["script"]
					if !ok {
						links = []string{}
					}
					linkMap["script"] = append(links, a.Val)
				}
			}
		case "img":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links, ok := linkMap["image"]
					if !ok {
						links = []string{}
					}
					linkMap["image"] = append(links, a.Val)
				}
			}
		}
	}
	if c := n.FirstChild; c != nil {
		linkMap = visit(linkMap, c)
	}
	if c := n.NextSibling; c != nil {
		linkMap = visit(linkMap, c)
	}
	return linkMap
}
