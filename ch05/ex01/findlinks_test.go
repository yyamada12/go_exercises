// Findlinks prints the links in an HTML document read from standard input.
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
	got := visit(nil, doc)

	// assert
	want := []string{"https://support.eji.org/give/153413/#!/donation/checkout", "/", "/doc/", "/pkg/", "/project/", "/help/", "/blog/", "https://play.golang.org/", "/dl/", "https://tour.golang.org/", "https://blog.golang.org/", "/doc/copyright.html", "/doc/tos.html", "http://www.google.com/intl/en/policies/privacy/", "http://golang.org/issues/new?title=x/website:", "https://google.com"}

	sort.Strings(got)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
