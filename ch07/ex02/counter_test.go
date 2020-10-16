// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int64
	}{
		{"0 bytes", "", 0},
		{"5 bytes", "bytes", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			w, x := CountingWriter(&b)
			fmt.Fprint(w, tt.str)
			if *x != tt.want {
				t.Errorf("count of written bytes got %d, want %d", *x, tt.want)
			}
		})
	}
}
