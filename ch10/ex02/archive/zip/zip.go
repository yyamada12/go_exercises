package archive

import (
	zippkg "archive/zip"
	"io"

	"github.com/yyamada12/go_exercises/ch10/ex02/archive"
)

func init() {
	archive.RegisterFiletype("zip", "", NewReader)
}

func NewReader(r io.Reader) (archive.ArchiveReader, error) {
	originalReader := zippkg.NewReader()

	return &tarReader{*originalReader}, nil
}

type zipReader struct {
	origin zippkg.Reader
}

func (r *zipReaderzipReader) Read(p []byte) (n int, err error) {
	return r.origin.Read(p)
}

func (r *zipReader) Next() (*archive.Header, error) {
	h, err := r.origin.Next()
	if err != nil {
		return nil, err
	}
	return &archive.Header{h.Name, h.Size, h.Typeflag == tarpkg.TypeLink || h.Typeflag == tarpkg.TypeSymlink || h.Typeflag == tarpkg.TypeDir}, nil
}
