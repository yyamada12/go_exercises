package archive

import (
	zippkg "archive/zip"
	"io"

	"github.com/yyamada12/go_exercises/ch10/ex02/archive"
)

func init() {
	archive.RegisterFiletype("zip", "PK", NewReader)
}

func NewReader(filename string) (archive.ArchiveReader, error) {
	originalReader, err := zippkg.OpenReader(filename)
	if err != nil {
		return nil, err
	}

	return &zipReader{origin: originalReader}, nil
}

type zipReader struct {
	origin *zippkg.ReadCloser
	reader *io.ReadCloser
	crt    int
}

func (r *zipReader) Read(p []byte) (n int, err error) {
	if r.reader != nil {
		reader := *r.reader
		return reader.Read(p)
	}
	reader, err := r.origin.File[r.crt].Open()
	if err != nil {
		return 0, err
	}
	r.reader = &reader
	return reader.Read(p)
}

func (r *zipReader) Next() (*archive.Header, error) {
	if r.reader != nil {
		reader := *r.reader
		reader.Close()
		r.reader = nil
	}
	r.crt++
	if r.crt >= len(r.origin.Reader.File) {
		return nil, io.EOF
	}
	f := r.origin.File[r.crt].FileInfo()
	return &archive.Header{f.Name(), f.Size(), f.IsDir()}, nil
}

func (r *zipReader) Close() error {
	return r.origin.Close()
}
