package main

import (
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func Test_visit(t *testing.T) {
	type want struct {
		element string
		count   int
	}

	// setup
	fp, err := os.Open("sample.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	doc, err := html.Parse(fp)
	if err != nil {
		log.Fatal(err)
	}

	// execute
	got := visit(map[string]int{}, doc)

	// assert
	wants := []want{{"html", 1}, {"title", 2}, {"a", 17}, {"p", 1}, {"div", 14}}

	for _, w := range wants {
		if count, ok := got[w.element]; !ok {
			t.Errorf("element %q not counted", w.element)
		} else if count != w.count {
			t.Errorf("count of element %q: got %d, want %d", w.element, count, w.count)
		}
	}

}
