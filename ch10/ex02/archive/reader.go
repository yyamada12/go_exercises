package archive

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

func main() {
	// r := NewReader("ch01.zip")

	// reader, format, err := NewReader("ch01.tar")
	// fmt.Println(format)
	// fmt.Println(err)
	// b, err := ioutil.ReadAll(reader)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(b)
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
	Close() error
}

type filetype struct {
	name, magic string
	newReader   func(string) (ArchiveReader, error)
}

var (
	filetypesMu     sync.Mutex
	atomicFiletypes atomic.Value
)

func RegisterFiletype(name, magic string, newReader func(string) (ArchiveReader, error)) {
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
func sniff(filename string) filetype {
	file, err := os.Open(filename)
	if err != nil {
		return filetype{}
	}
	defer file.Close()

	r := bufio.NewReader(file)
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
func NewReader(filename string) (ArchiveReader, string, error) {
	f := sniff(filename)
	if f.newReader == nil {
		return nil, "", fmt.Errorf("archive: unknown format")
	}
	m, err := f.newReader(filename)
	return m, f.name, err
}
