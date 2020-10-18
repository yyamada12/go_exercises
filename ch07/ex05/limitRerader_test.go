package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestLimitReader(t *testing.T) {

	type args struct {
		r io.Reader
		n int64
	}
	tests := []struct {
		name  string
		input string
		n     int64
		want  string
	}{
		{"n is grater than string length", "long string", 7, "long st"},
		{"n is less than string length", "short", 7, "short"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bytes.NewBufferString(tt.input)
			lr := LimitReader(r, tt.n)
			got, err := ioutil.ReadAll(lr)
			if err != nil {
				t.Errorf("error at Read() from LimitReader: %s", err.Error())
			}

			if string(got) != tt.want {
				t.Errorf("Read() from LimitReader got %q, want %q", got, tt.want)
			}
		})
	}
}
