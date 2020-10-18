package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	r := bytes.NewBufferString("hello world")
	lr := LimitReader(r, 8)
	p := make([]byte, 6)
	q := make([]byte, 6)
	n, _ := lr.Read(p)
	fmt.Printf("%d byte read: %s\n", n, p)
	m, _ := lr.Read(q)
	fmt.Printf("%d byte read: %s\n", m, q)
}

// LimitReader returns a io.Reader that return EOF after read n bytes
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{n, r}
}

type limitReader struct {
	n      int64
	origin io.Reader
}

func (w *limitReader) Read(p []byte) (int, error) {
	if w.n == 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > w.n {
		p = p[0:w.n]
	}
	n, err := w.origin.Read(p)
	w.n -= int64(n)
	return n, err
}
