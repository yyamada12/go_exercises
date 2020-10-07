// Prettyprint prints HTML document in pretty format.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout
var depth int
var escape bool
var emptyElements map[string]bool

func init() {
	emptyElements = map[string]bool{}
	for _, key := range []string{"br", "hr", "img", "input", "meta", "area", "base", "col", "embed", "keygen", "link", "param", "source"} {
		emptyElements[key] = true
	}
}

func main() {
	for _, url := range os.Args[1:] {
		prettyprint(url)
	}
}

func prettyprint(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		handleElementNode(n)
	} else if n.Type == html.CommentNode || n.Type == html.TextNode {
		handleTextNode(n)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil && emptyElements[n.Data] {
			// shortcut close tag
			return
		}
		depth--
		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
	}
}

func handleElementNode(n *html.Node) {
	var attrs strings.Builder
	for _, a := range n.Attr {
		attrs.WriteString(" ")
		attrs.WriteString(a.Key)
		if a.Val != "" {
			attrs.WriteString(`="`)
			attrs.WriteString(a.Val)
			attrs.WriteString(`"`)
		}
	}

	if n.FirstChild == nil && emptyElements[n.Data] {
		// shortcut close tag
		fmt.Fprintf(out, "%*s<%s%s/>\n", depth*2, "", n.Data, attrs.String())
		return
	}
	fmt.Fprintf(out, "%*s<%s%s>\n", depth*2, "", n.Data, attrs.String())
	depth++
}

func handleTextNode(n *html.Node) {
	for _, line := range strings.Split(n.Data, "\n") {
		text := strings.TrimSpace(line)
		if escape {
			text = html.EscapeString(text)
		}
		if text != "" {
			fmt.Fprintf(out, "%*s%s\n", (depth+1)*2, "", text)
		}
	}
}
