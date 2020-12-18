// Charcount computes counts of Unicode characters.
package main

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func Test_charcount(t *testing.T) {
	tests := []struct {
		name  string
		arg   string
		want  map[rune]int
		want1 [utf8.UTFMax + 1]int
		want2 int
	}{
		{"empty string", "", map[rune]int{}, [utf8.UTFMax + 1]int{}, 0},
		{"ascii string", "ascii", map[rune]int{'a': 1, 's': 1, 'c': 1, 'i': 2}, [utf8.UTFMax + 1]int{1: 5}, 0},
		{"multi byte string", "¥あ👀", map[rune]int{'¥': 1, 'あ': 1, '👀': 1}, [utf8.UTFMax + 1]int{2: 1, 3: 1, 4: 1}, 0},
		{"invalid string", "invalid\xf0\x28", map[rune]int{'i': 2, 'n': 1, 'v': 1, 'a': 1, 'l': 1, 'd': 1, '\x28': 1}, [utf8.UTFMax + 1]int{1: 8}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := charcount(strings.NewReader(tt.arg))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("charcount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("charcount() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("charcount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
