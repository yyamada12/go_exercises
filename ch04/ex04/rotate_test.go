package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"n=0,len=5", args{[]int{1, 2, 3, 4, 5}, 0}, []int{1, 2, 3, 4, 5}},
		{"n=2,len=5", args{[]int{1, 2, 3, 4, 5}, 2}, []int{3, 4, 5, 1, 2}},
		{"n=5,len=5", args{[]int{1, 2, 3, 4, 5}, 5}, []int{1, 2, 3, 4, 5}},
		{"n=8,len=5", args{[]int{1, 2, 3, 4, 5}, 8}, []int{4, 5, 1, 2, 3}},
		{"n=-3,len=5", args{[]int{1, 2, 3, 4, 5}, -3}, []int{3, 4, 5, 1, 2}},
		{"n=3,len=6", args{[]int{1, 2, 3, 4, 5, 6}, 3}, []int{4, 5, 6, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			leftRotate(tt.args.s, tt.args.n)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("rotate %d, got %q, want %q", tt.args.n, tt.args.s, tt.want)
			}
		})
	}
}
