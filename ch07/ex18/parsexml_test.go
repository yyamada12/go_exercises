package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_ParseXML(t *testing.T) {

	// setup
	fp, err := os.Open("sample.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	sample, err := ioutil.ReadFile("sample.xml")
	if err != nil {
		log.Fatal(err)
	}

	// execute
	root, err := ParseXML(fp)
	if err != nil {
		t.Errorf("ParseXML got error, %s", err.Error())
	}

	// assert
	got := root.String()
	want := string(strings.TrimRight(string(sample), "\n"))
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
