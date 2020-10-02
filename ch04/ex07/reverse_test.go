package main

import (
	"bytes"
	"testing"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want []byte
	}{
		{"only ASCII", []byte("hoge 123_"), []byte("_321 egoh")},
		{"only 3-byte", []byte("ほげほげ"), []byte("げほげほ")},
		{"mix 1-byte to 4-byte", []byte("a©世🤔"), []byte("🤔世©a")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.arg)
			if !bytes.Equal(tt.arg, tt.want) {
				t.Errorf("got %q, want %q", tt.arg, tt.want)
			}
		})
	}
}
