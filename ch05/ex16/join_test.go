package main

import "testing"

func Test_join(t *testing.T) {
	type args struct {
		sep   string
		elems []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"no elems", args{",", []string{}}, ""},
		{"1 elem", args{",", []string{"hoge"}}, "hoge"},
		{"2 elems", args{",", []string{"hoge", "fuga"}}, "hoge,fuga"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := join(tt.args.sep, tt.args.elems...); got != tt.want {
				t.Errorf("join(%q, %q...) = %v, want %v", tt.args.sep, tt.args.elems, got, tt.want)
			}
		})
	}
}
