// The du command computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type result struct {
	root   string
	nfiles int64
	nbytes int64
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	num := len(roots)

	// Traverse each root of the file tree in parallel.
	fileSizes := make([]chan int64, num)
	for i := range fileSizes {
		fileSizes[i] = make(chan int64)
	}
	for i, root := range roots {
		go func(dir string, fileSize chan<- int64) {
			var wg sync.WaitGroup
			wg.Add(1)
			go walkDir(dir, &wg, fileSize)
			wg.Wait()
			close(fileSize)
		}(root, fileSizes[i])
	}

	results := make(chan result)
	for i, root := range roots {
		go aggregate(root, fileSizes[i], results)
	}

	for i := 0; i < num; i++ {
		res := <-results
		printDiskUsage(res)
	}
}

func printDiskUsage(res result) {
	fmt.Printf("%s: %d files  %.1f GB\n", res.root, res.nfiles, float64(res.nbytes)/1e9)
}

func aggregate(root string, fileSize <-chan int64, results chan<- result) {
	var nfiles, nbytes int64
	for size := range fileSize {
		nfiles++
		nbytes += size
	}
	results <- result{root, nfiles, nbytes}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, wg *sync.WaitGroup, fileSize chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, wg, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
