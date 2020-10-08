// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println(`USAGE:
	go run elementByID.go [URL] [id]`)
	}

	url, id := os.Args[1], os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	element := ElementByID(doc, id)
	if element != nil {
		fmt.Println("element", element.Data, "found")
	} else {
		fmt.Println("element not found")
	}
}

// ElementByID finds element from HTML document by id
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, findElementByIDFunc(id), doNothing)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		stop := pre(n)
		if stop {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil {
		stop := post(n)
		if stop {
			return n
		}
	}

	return nil
}

func findElementByIDFunc(id string) func(n *html.Node) bool {
	return func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					return true
				}
			}
		}
		return false
	}
}

func doNothing(n *html.Node) bool {
	return false
}
