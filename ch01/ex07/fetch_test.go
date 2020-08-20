package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	// setup
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	out = new(bytes.Buffer)

	// excercise
	fetch(ts.URL)

	// verify
	want := "Hello, client\n"
	got := out.(*bytes.Buffer).String()
	if got != want {
		t.Errorf("fetch(%s) = %q, want \"%s\"", ts.URL, got, want)
	}
}
