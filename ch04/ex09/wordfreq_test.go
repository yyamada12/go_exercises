package main

import (
	"strings"
	"testing"
)

func Test_wordfreq(t *testing.T) {
	type want struct {
		key string
		cnt int
	}

	// setup
	text := "hoge hoge fuga"
	wants := []want{
		{"hoge", 2},
		{"fuga", 1},
	}

	// exec
	got := wordfreq(strings.NewReader(text))

	// assert
	for _, w := range wants {
		if cnt, ok := got[w.key]; !ok {
			t.Errorf("counts not contain key %q", w.key)
		} else if cnt != w.cnt {
			t.Errorf("count of %q got %d, want %d", w.key, cnt, w.cnt)
		}
	}
}
