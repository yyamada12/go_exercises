package archive

import (
	"archive/zip"
	"bufio"
	"io"
	"sync"
	"sync/atomic"
)

func main() {
	// r := NewReader("ch01.zip")

	reader, _ := zip.OpenReader("ch01.zip")
	reader.File[0].FileInfo().IsDir()
}

type Header struct {
	Name  string // Name of file entry
	Size  int64  // Logical file size in bytes
	IsDir bool
}

type Reader struct {
	r io.Reader
}

type ArchiveReader interface {
	Next() (*Header, error)
	io.Reader
}

type filetype struct {
	name, magic string
	newReader   func(io.Reader) (ArchiveReader, error)
}

var (
	filetypesMu     sync.Mutex
	atomicFiletypes atomic.Value
)

func RegisterFiletype(name, magic string, newReader func(io.Reader) (ArchiveReader, error)) {
	filetypesMu.Lock()
	filetypes, _ := atomicFiletypes.Load().([]filetype)
	atomicFiletypes.Store(append(filetypes, filetype{name, magic, newReader}))
	filetypesMu.Unlock()
}

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of r's data.
func sniff(r reader) filetype {
	filetypes, _ := atomicFiletypes.Load().([]filetype)
	for _, f := range filetypes {
		b, err := r.Peek(len(f.magic))
		if err == nil && match(f.magic, b) {
			return f
		}
	}
	return filetype{}
}

// NewReader read a file that has been archived in a registered filetype.
// The string returned is the filetype name used during filetype registration.
// Filetype registration is typically done by an init function in the archive-
// specific package.
func NewReader(r io.Reader) (Reader, string, error) {
	rr := asReader(r)
	f := sniff(rr)
	if f.newReader == nil {
		return nil, "", ErrFormat
	}
	m, err := f.newReader(rr)
	return m, f.name, err
}
