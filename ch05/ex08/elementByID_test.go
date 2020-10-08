// Findlinks prints the links in an HTML document read from standard input.
package main

import (
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func Test_ElementByID(t *testing.T) {
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

	t.Run("not exist element", func(t *testing.T) {
		// execute
		got := ElementByID(doc, "noElement")

		// asert
		if got != nil {
			t.Errorf("want nil, got %q", got.Data)
		}
	})

	t.Run("exist element", func(t *testing.T) {
		// execute
		got := ElementByID(doc, "page")

		// asert
		if got.Type != html.ElementNode {
			t.Errorf("want element, got %q", got.Type)
		} else if got.Data != "main" {
			t.Errorf("want main, got %q", got.Data)
		}
	})
}
