package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	// setup
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	tests := []struct {
		name string
		url  string
	}{
		{"normal request", ts.URL},
		{"without 'http://'", strings.Replace(ts.URL, "http://", "", 1)},
	}
	want := "Hello, client\n"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// excercise
			out = new(bytes.Buffer)
			fetch(tt.url)

			// verify
			got := out.(*bytes.Buffer).String()
			if got != want {
				t.Errorf("fetch(%s) = %q, want %q", tt.url, got, want)
			}
		})
	}
}
