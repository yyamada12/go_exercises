package main

import (
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func Test_ElementsByTagName(t *testing.T) {
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
	got := ElementsByTagName(doc, "img")

	// assert
	want := []string{"Header-logo", "HeroDownloadButton-image", "Footer-gopher"}
	if len(got) != len(want) {
		t.Errorf("got %d nodes, want %d nodes", len(got), len(want))
	}

	for i, g := range got {
		if g.Data != "img" {
			t.Errorf("got %s, want img", g.Data)
		}
		for _, a := range g.Attr {
			if a.Key == "class" && a.Val != want[i] {
				t.Errorf("got class=%q, want class=%q", a.Val, want[i])
			}
		}
	}

}
