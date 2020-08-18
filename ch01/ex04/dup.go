package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	dup(os.Args[1:])
}

func dup(files []string) {
	counts := make(map[string]int)
	filenamesByLine := make(map[string]map[string]struct{})
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts, filenamesByLine)
	} else {
		for _, file := range files {
			processFile(file, counts, filenamesByLine)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
			for filename := range filenamesByLine[line] {
				fmt.Printf("\t%s", filename)
			}
			fmt.Printf("\n")
		}

	}
}

func processFile(filename string, counts map[string]int, filenamesByLine map[string]map[string]struct{}) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		return
	}
	defer f.Close()
	countLines(f, filename, counts, filenamesByLine)
}

func countLines(f io.Reader, filename string, counts map[string]int, filenamesByLine map[string]map[string]struct{}) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++

		// initialize set(map[string]struct{}) for the line
		if _, ok := filenamesByLine[line]; !ok {
			filenamesByLine[line] = make(map[string]struct{})
		}
		// add filename to the set
		filenamesByLine[line][filename] = struct{}{}

	}
	// NOTE: ignoring potential errors from input.Err()
}
