package main

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"angram", args{"anagram", "managar"}, true},
		{"not angram", args{"anagram", "manager"}, false},
		{"same string", args{"anagram", "anagram"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagram(%s, %s) = %v, want %v", tt.args.s1, tt.args.s2, got, tt.want)
			}
		})
	}
}
