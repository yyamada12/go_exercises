// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", stackString(stack), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].Name.Local == y[0] {
			y = y[1:]
		} else {
			for _, attr := range x[0].Attr {
				if attr.Value == y[0] {
					y = y[1:]
					break
				}
			}
		}
		x = x[1:]
	}
	return false
}

func stackString(x []xml.StartElement) string {
	s := strings.Builder{}
	for i, e := range x {
		if i != 0 {
			s.WriteRune(' ')
		}
		s.WriteString(e.Name.Local)
	}
	return s.String()
}
