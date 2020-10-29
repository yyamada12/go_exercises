package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Node is XML node (one of CharData, Element)
type Node interface {
	fmt.Stringer
}

// CharData is char data
type CharData string

// Element is an element node
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (c CharData) String() string {
	return string(c)
}

func (e *Element) String() string {
	var s strings.Builder
	s.WriteByte('<')
	s.WriteString(stringify(e.Type))
	for _, attr := range e.Attr {
		s.WriteByte(' ')
		s.WriteString(stringify(attr.Name))
		s.WriteByte('=')
		s.WriteString(attr.Value)
	}
	if len(e.Children) == 0 {
		s.WriteString("/>")
		return s.String()
	}
	s.WriteByte('>')
	for _, c := range e.Children {
		s.WriteString(c.String())
	}
	s.WriteString("</")
	s.WriteString(stringify(e.Type))
	s.WriteByte('>')
	return s.String()
}

func stringify(n xml.Name) string {
	if n.Space != "" {
		return n.Space + ":" + n.Local
	}
	return n.Local
}

func main() {
	root, err := ParseXML(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(root)
}

// ParseXML parses a xml document and returns xml node tree
func ParseXML(in io.Reader) (Node, error) {
	dec := xml.NewDecoder(in)
	var stack []*Element // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return nil, fmt.Errorf("unexpected eof")
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			e := Element{Type: tok.Name, Attr: tok.Attr}
			if len(stack) >= 1 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, &e)
			}
			stack = append(stack, &e) // push
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("end element comes before start element")
			} else if len(stack) == 1 {
				return stack[0], nil
			}
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) >= 1 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(tok))
			} else {
				return nil, fmt.Errorf("char data comes before start element")
			}
		}
	}
}
