package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_echo(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{"no lines", args{[]string{}}, ""},
		{"single line", args{[]string{"first line"}}, "0 first line\n"},
		{"two lines", args{[]string{"first line", "second line"}}, "0 first line\n1 second line\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			descr := fmt.Sprintf("echo(%q)", tt.args.args)
			out = new(bytes.Buffer)
			echo(tt.args.args)
			got := out.(*bytes.Buffer).String()
			if got != tt.want {
				t.Errorf("%s = %q, want %q", descr, got, tt.want)
			}

		})
	}
}
