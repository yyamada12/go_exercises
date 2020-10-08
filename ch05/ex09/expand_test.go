package main

import "testing"

func Test_expand(t *testing.T) {
	type args struct {
		s string
		f func(string) string
	}
	tests := []struct {
		name     string
		args     args
		funcName string
		want     string
	}{
		{"echo", args{"$echo1, $echo2", echo}, "echo", "echo1, echo2"},
		{"reverse", args{"$rev1, $$re$v2", reverse}, "reverse", "1ver, $er2v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("expand(%s, %s) = %v, want %v", tt.args.s, tt.funcName, got, tt.want)
			}
		})
	}
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
