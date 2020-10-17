package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func TestParse(t *testing.T) {
	// setup
	fp, err := os.Open("sample.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	want, err := html.Parse(fp)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := ioutil.ReadFile("sample.html")
	if err != nil {
		log.Fatal(err)
	}
	arg := string(doc)

	// execute
	got, err := Parse(arg)

	// assert
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() got %v, want %v", got, want)
	}

}
