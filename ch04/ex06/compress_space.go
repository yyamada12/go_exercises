package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println(string(compressSpace([]byte("hello　世界!"))))
}

func compressSpace(s []byte) []byte {
	i := 0
	for j := 0; j < len(s); {
		r, size := utf8.DecodeRune(s[j:])
		if unicode.IsSpace(r) {
			s[i] = ' '
			i++
		} else {
			copy(s[i:i+size], s[j:j+size])
			i += size
		}
		j += size
	}
	return s[:i]
}
