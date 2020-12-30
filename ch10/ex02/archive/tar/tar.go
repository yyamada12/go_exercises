package archive

import (
	tarpkg "archive/tar"
	"bufio"
	"os"
	"strings"

	"github.com/yyamada12/go_exercises/ch10/ex02/archive"
)

func init() {
	archive.RegisterFiletype("tar", strings.Repeat("?", 257)+"ustar", NewReader)
}

func NewReader(filename string) (archive.ArchiveReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(file)

	originalReader := tarpkg.NewReader(r)
	return &tarReader{originalReader, file}, nil
}

type tarReader struct {
	origin *tarpkg.Reader
	file   *os.File
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

func (r *tarReader) Close() error {
	return r.file.Close()
}
