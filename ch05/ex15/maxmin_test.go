package main

import (
	"testing"
)

func Test_max1(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1 number", []int{1}, 1},
		{"2 numbers", []int{1, 2}, 2},
		{"3 numbers", []int{1, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max1(tt.nums...); got != tt.want {
				t.Errorf("max1(%d...) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func Test_max2(t *testing.T) {

	tests := []struct {
		name string
		n    int
		nums []int
		want int
	}{
		{"1 number", 1, []int{}, 1},
		{"2 numbers", 1, []int{2}, 2},
		{"3 numbers", 1, []int{2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max2(tt.n, tt.nums...); got != tt.want {
				t.Errorf("max1(%d, %d...) = %v, want %v", tt.n, tt.nums, got, tt.want)
			}
		})
	}
}

func Test_min1(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1 number", []int{1}, 1},
		{"2 numbers", []int{1, 2}, 1},
		{"3 numbers", []int{1, 2, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min1(tt.nums...); got != tt.want {
				t.Errorf("min1(%d...) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func Test_min2(t *testing.T) {

	tests := []struct {
		name string
		n    int
		nums []int
		want int
	}{
		{"1 number", 1, []int{}, 1},
		{"2 numbers", 1, []int{2}, 1},
		{"3 numbers", 1, []int{2, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min2(tt.n, tt.nums...); got != tt.want {
				t.Errorf("max1(%d, %d...) = %v, want %v", tt.n, tt.nums, got, tt.want)
			}
		})
	}
}
