// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import "testing"

func TestWordCounter_Write(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want WordCounter
	}{
		{"no word", []byte(""), WordCounter(0)},
		{"1 word", []byte("word"), WordCounter(1)},
		{"2 word", []byte("2    \tword"), WordCounter(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c WordCounter
			c.Write(tt.arg)
			if c != tt.want {
				t.Errorf("WordCounter.Write(%q) = %v, want %v", tt.arg, c, tt.want)
			}
		})
	}
}

func TestLineCounter_Write(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want LineCounter
	}{
		{"no line", []byte(""), LineCounter(0)},
		{"1 line", []byte("line"), LineCounter(1)},
		{"2 line", []byte("2    \nline"), LineCounter(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c LineCounter
			c.Write(tt.arg)
			if c != tt.want {
				t.Errorf("LineCounter.Write(%q) = %v, want %v", tt.arg, c, tt.want)
			}
		})
	}
}
