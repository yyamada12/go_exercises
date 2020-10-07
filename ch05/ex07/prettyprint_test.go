// Findlinks prints the links in an HTML document read from standard input.
package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func Test_prettyprintNoEscape(t *testing.T) {
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
	out = new(bytes.Buffer)
	forEachNode(doc, startElement, endElement)
	got := out.(*bytes.Buffer)

	// assert
	if _, err := html.Parse(got); err != nil {
		t.Errorf("got parse err: %q", err)
	}
}

func Test_prettyprintEscape(t *testing.T) {
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
	escape = true
	out = new(bytes.Buffer)
	forEachNode(doc, startElement, endElement)
	got := out.(*bytes.Buffer)

	// assert
	if err := validHTML(got); err != nil {
		t.Errorf("got parse err: %q", err)
	}
}

func validHTML(r io.Reader) error {
	d := xml.NewDecoder(r)

	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	for {
		_, err := d.Token()
		switch err {
		case io.EOF:
			return nil
		case nil:
		default:
			return err
		}
	}
}
