// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import "testing"

func Test_tree_string(t *testing.T) {

	tests := []struct {
		name   string
		values []int
		want   string
	}{
		{"no elements", []int{}, "[]"},
		{"1 element", []int{1}, "[1]"},
		{"2 elements", []int{1, 2}, "[1 2]"},
		{"sorted 5 elements", []int{5, 2, 3, 4, 1}, "[1 2 3 4 5]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *tree
			for _, v := range tt.values {
				root = add(root, v)
			}
			if got := root.string(); got != tt.want {
				t.Errorf("tree.string() = %v, want %v", got, tt.want)
			}
		})
	}
}
