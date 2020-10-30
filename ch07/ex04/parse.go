package main

import (
	"io"
	"io/ioutil"
	"log"

	"golang.org/x/net/html"
)

func main() {
	doc, err := ioutil.ReadFile("sample.html")
	if err != nil {
		log.Fatal(err)
	}
	arg := string(doc)

	_, err = Parse(arg)
	if err != nil {
		log.Fatal(err)
	}
}

func Parse(s string) (*html.Node, error) {
	r := NewReader(s)
	return html.Parse(r)
}

func NewReader(s string) io.Reader {
	var r stringReader
	r.b = []byte(s)
	return &r
}

type stringReader struct {
	b []byte
}

func (s *stringReader) Read(p []byte) (n int, err error) {
	if len(s.b) == 0 {
		return 0, io.EOF
	}
	n = copy(p, s.b)
	s.b = s.b[n:]
	return n, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
