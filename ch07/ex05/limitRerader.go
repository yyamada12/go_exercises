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
	return &limitReader{r, n}
}

type limitReader struct {
	origin io.Reader
	n      int64
}

func (r *limitReader) Read(p []byte) (int, error) {
	if r.n == 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.n {
		p = p[0:r.n]
	}
	n, err := r.origin.Read(p)
	r.n -= int64(n)
	return n, err
}
