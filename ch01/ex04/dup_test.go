package main

import (
	"bytes"
	"testing"
)

func Test_countLines(t *testing.T) {

	counts := make(map[string]int)
	filenamesByLine := make(map[string]map[string]struct{})

	f1 := bytes.NewBufferString("foo\nbar\nfoo")
	f2 := bytes.NewBufferString("bar\nbaz")

	countLines(f1, "file1", counts, filenamesByLine)
	countLines(f2, "file2", counts, filenamesByLine)

	if counts["foo"] != 2 {
		t.Errorf("counts[\"foo\"] = %d, want 2", counts["foo"])
	}
	if counts["bar"] != 2 {
		t.Errorf("counts[\"bar\"] = %d, want 2", counts["bar"])
	}

	if _, ok := filenamesByLine["foo"]["file1"]; !ok {
		t.Errorf("filenamesByLine[\"foo\"] not contain \"file1\"")
	}

	if _, ok := filenamesByLine["bar"]["file1"]; !ok {
		t.Errorf("filenamesByLine[\"bar\"] not contain \"file1\"")
	}

	if _, ok := filenamesByLine["bar"]["file2"]; !ok {
		t.Errorf("filenamesByLine[\"bar\"] not contain \"file2\"")
	}
}
