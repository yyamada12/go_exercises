// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	invalid := 0                   // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.Is(unicode.L, r) {
			counts["Letter"]++
		} else if unicode.Is(unicode.M, r) {
			counts["mark"]++
		} else if unicode.Is(unicode.N, r) {
			counts["number"]++
		} else if unicode.Is(unicode.C, r) {
			counts["control"]++
		} else if unicode.Is(unicode.P, r) {
			counts["punctuation"]++
		} else if unicode.Is(unicode.S, r) {
			counts["symbol"]++
		} else if unicode.Is(unicode.Z, r) {
			counts["space"]++
		} else {
			counts["other"]++
		}

	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
