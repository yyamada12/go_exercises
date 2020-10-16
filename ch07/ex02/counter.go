package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var b bytes.Buffer
	w, x := CountingWriter(&b)
	fmt.Fprintf(w, "%s", "hoge")
	fmt.Println(*x)
}

// CountingWriter returns Writer wrapping given Writer for counting and address of counter
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var x countingWriter
	x.origin = w
	return &x, &x.count
}

type countingWriter struct {
	count  int64
	origin io.Writer
}

func (w *countingWriter) Write(p []byte) (int, error) {
	w.count += int64(len(p)) // convert int to wrapperWriter
	return w.origin.Write(p)
}
