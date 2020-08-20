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
		if r.URL.Path == "/err" {
			http.Error(w, "Not Found.", http.StatusNotFound)
			return
		}
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	tests := []struct {
		name string
		url  string
		want string
	}{
		{"normal request", ts.URL, "200 OK\nHello, client\n"},
		{"without 'http://'", strings.Replace(ts.URL, "http://", "", 1), "200 OK\nHello, client\n"},
		{"not found", ts.URL + "/err", "404 Not Found\nNot Found.\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// excercise
			out = new(bytes.Buffer)
			fetch(tt.url)

			// verify
			got := out.(*bytes.Buffer).String()
			if got != tt.want {
				t.Errorf("fetch(%s) = %q, want %q", tt.url, got, tt.want)
			}
		})
	}
}
