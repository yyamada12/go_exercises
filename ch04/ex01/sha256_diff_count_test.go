package main

import (
	"testing"
)

func Test_diffCount(t *testing.T) {
	type args struct {
		d1 [32]byte
		d2 [32]byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			"same hash",
			args{
				[32]byte{0xd, 0xe, 0xa, 0xd, 0xb, 0xe, 0xe, 0xf},
				[32]byte{0xd, 0xe, 0xa, 0xd, 0xb, 0xe, 0xe, 0xf},
			},
			byte(0),
		},
		{
			"1bit diff",
			args{
				[32]byte{0xd, 0xe, 0xa, 0xd, 0xb, 0xe, 0xe, 0xf, 0},
				[32]byte{0xd, 0xe, 0xa, 0xd, 0xb, 0xe, 0xe, 0xf, 1},
			},
			byte(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffCount(tt.args.d1, tt.args.d2); got != tt.want {
				t.Errorf("diffCount(%q, %q) = %v, want %v", tt.args.d1, tt.args.d2, got, tt.want)
			}
		})
	}
}
