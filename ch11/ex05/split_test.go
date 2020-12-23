package main

import (
	"strings"
	"testing"
)

func TestSplt(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty string sep by empty string", args{"", ""}, 0},
		{"some string sep by empty string", args{"test", ""}, 4},
		{"empty string sep by some string", args{"", ","}, 1},
		{"3 elements sep by colon", args{"a:b:c", ":"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			words := strings.Split(tt.args.s, tt.args.sep)
			if got := len(words); got != tt.want {
				t.Errorf("Split(%q, %q) returned %d words, want %d", tt.args.s, tt.args.sep, got, tt.want)
			}
		})
	}
}
