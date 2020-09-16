package main

import "testing"

func Test_delimiter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args string
		want string
	}{
		{"1 Digit Int", "1", "1"},
		{"2 Digit Int", "12", "12"},
		{"3 Digit Int", "123", "123"},
		{"4 Digit Int", "1234", "1,234"},
		{"10 Digit Int", "1234567890", "1,234,567,890"},
		{"10 Digit Float", "1234.567890", "1,234.567 890"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delimiter(tt.args); got != tt.want {
				t.Errorf("delimiter(%s) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
