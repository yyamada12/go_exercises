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
		{"empty", args{[]string{}}, "\n"},
		{"single", args{[]string{"first"}}, "first\n"},
		{"second", args{[]string{"first", "second"}}, "first second\n"},
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
