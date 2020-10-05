package main

import (
	"log"
	"os"
	"reflect"
	"sort"
	"testing"

	"golang.org/x/net/html"
)

func Test_visit(t *testing.T) {
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
	got := visit(map[string][]string{}, doc)

	// assert
	want := map[string][]string{
		"hyperlink": {
			"https://support.eji.org/give/153413/#!/donation/checkout",
			"/",
			"/doc/",
			"/pkg/",
			"/project/",
			"/help/",
			"/blog/",
			"https://play.golang.org/",
			"/dl/",
			"https://tour.golang.org/",
			"https://blog.golang.org/",
			"/doc/copyright.html",
			"/doc/tos.html",
			"http://www.google.com/intl/en/policies/privacy/",
			"http://golang.org/issues/new?title=x/website:",
			"https://google.com",
		},
		"stylesheet": {
			"https://fonts.googleapis.com/css?family=Work+Sans:600|Roboto:400,700",
			"https://fonts.googleapis.com/css?family=Product+Sans&text=Supported%20by%20Google&display=swap",
			"/lib/godoc/style.css",
		},
		"script": {
			"/lib/godoc/jquery.js",
			"/lib/godoc/playground.js",
			"/lib/godoc/godocs.js",
		},
		"image": {
			"/lib/godoc/images/go-logo-blue.svg",
			"/lib/godoc/images/cloud-download.svg",
			"/lib/godoc/images/footer-gopher.jpg",
		},
	}

	for _, key := range []string{"hyperlink", "stylesheet", "script", "image"} {
		g, ok := got[key]
		if !ok {
			t.Errorf("key \"hyperlink\" has no item")
		}
		w, _ := want[key]
		sort.Strings(g)
		sort.Strings(w)
		if !reflect.DeepEqual(g, w) {
			t.Errorf("in %q key got %v,want %v", key, g, w)
		}
	}
}
