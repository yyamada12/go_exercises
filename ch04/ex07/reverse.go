// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte("hello 世界!")
	reverseString(s)
	fmt.Println(string(s))
}

func reverseString(s []byte) {
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		reverse(s[i : i+size])
		i += size
	}
	reverse(s)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
