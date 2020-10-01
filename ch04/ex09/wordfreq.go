package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := wordfreq(os.Stdin)
	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

func wordfreq(input io.Reader) map[string]int {
	counts := make(map[string]int) // counts of words
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}
	return counts
}
