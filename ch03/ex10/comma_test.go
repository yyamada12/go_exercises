package main

import "testing"

func Test_comma(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"1 Digit", "1", "1"},
		{"2 Digit", "12", "12"},
		{"3 Digit", "123", "123"},
		{"4 Digit", "1234", "1,234"},
		{"10 Digit", "1234567890", "1,234,567,890"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.args); got != tt.want {
				t.Errorf("comma(%s) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
