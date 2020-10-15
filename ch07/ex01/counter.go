package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter demonstrates an implementation of io.Writer that counts words in p.
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	in := bufio.NewScanner(bytes.NewReader(p))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		*c += WordCounter(1)
	}
	return len(p), nil
}

// LineCounter demonstrates an implementation of io.Writer that counts words in p.
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	in := bufio.NewScanner(bytes.NewReader(p))
	in.Split(bufio.ScanLines)
	for in.Scan() {
		*c += LineCounter(1)
	}
	return len(p), nil
}

func main() {
	var c WordCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "1"

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "2"
}
