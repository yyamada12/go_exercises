// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

import (
	"io"
	"os/exec"
)

type writer struct {
	in  io.WriteCloser
	cmd *exec.Cmd
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) (io.WriteCloser, error) {
	cmd := exec.Command("bzip2")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	cmd.Stdout = out
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return &writer{stdin, cmd}, nil
}

func (w *writer) Write(data []byte) (int, error) {
	return w.in.Write(data)
}

// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	w.in.Close()
	return w.cmd.Wait()
}
