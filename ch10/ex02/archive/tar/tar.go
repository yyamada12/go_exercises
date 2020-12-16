package archive

import (
	tarpkg "archive/tar"
	"io"

	"github.com/yyamada12/go_exercises/ch10/ex02/archive"
)

func init() {
	archive.RegisterFiletype("tar", "", NewReader)
}

func NewReader(r io.Reader) (archive.ArchiveReader, error) {
	originalReader := tarpkg.NewReader(r)

	return &tarReader{*originalReader}, nil
}

type tarReader struct {
	origin tarpkg.Reader
}

func (r *tarReader) Read(p []byte) (n int, err error) {
	return r.origin.Read(p)
}

func (r *tarReader) Next() (*archive.Header, error) {
	h, err := r.origin.Next()
	if err != nil {
		return nil, err
	}
	return &archive.Header{h.Name, h.Size, h.Typeflag == tarpkg.TypeLink || h.Typeflag == tarpkg.TypeSymlink || h.Typeflag == tarpkg.TypeDir}, nil
}
