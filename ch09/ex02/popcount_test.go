package popcount

import "testing"

func TestPopCount(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"min value", args{0}, 0},
		{"0xffff", args{0xffff}, 16},
		{"max value", args{18446744073709551615}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCount(tt.args.x); got != tt.want {
				t.Errorf("PopCount(%d) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
