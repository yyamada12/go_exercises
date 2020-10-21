package main

import (
	"sort"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args sort.Interface
		want bool
	}{
		{"int slice", sort.IntSlice([]int{1, 3, 4, 3, 1}), true},
		{"odd length string", sortableString("level"), true},
		{"even length string", sortableString("noon"), true},
		{"empty string", sortableString(""), true},
		{"not palindrome string", sortableString("palindrome"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args); got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
