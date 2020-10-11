package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	fp, err := os.Open("sample.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	doc, err := html.Parse(fp)
	if err != nil {
		log.Fatal(err)
	}
	images := ElementsByTagName(doc, "img")
	fmt.Println("<images>")
	for _, image := range images {
		for _, a := range image.Attr {
			fmt.Println(a.Key, a.Val)
		}
		fmt.Println()
	}

	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Println("<headings>")
	for _, heading := range headings {
		for _, a := range heading.Attr {
			fmt.Println(a.Key, a.Val)
		}
		fmt.Println()
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	names = name
	return visit(nil, doc)
}

var names []string

func visit(elements []*html.Node, n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && contains(names, n.Data) {
		elements = append(elements, n)
	}
	if c := n.FirstChild; c != nil {
		elements = visit(elements, c)
	}
	if c := n.NextSibling; c != nil {
		elements = visit(elements, c)
	}
	return elements
}

func contains(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}
