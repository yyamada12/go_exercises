package popcount

import "testing"

func TestPopCountBitShift(t *testing.T) {
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
			if got := PopCountBitShift(tt.args.x); got != tt.want {
				t.Errorf("PopCount(%d) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountBitShift(0x1234567890ABCDEF)
	}
}
