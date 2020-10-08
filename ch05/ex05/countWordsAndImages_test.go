package main

import (
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func Test_countWordsAndImages(t *testing.T) {
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
	gotWords, gotImages := countWordsAndImages(doc)

	// assert
	wantWords, wantImages := 123, 3
	if gotWords != wantWords {
		t.Errorf("count of words got %d, wnat %d", gotWords, wantWords)
	}

	if gotImages != wantImages {
		t.Errorf("count of images got %d, wnat %d", gotImages, wantImages)
	}

}
