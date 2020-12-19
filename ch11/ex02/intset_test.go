package intset

import (
	"testing"
)

func TestIntSet_Has(t *testing.T) {
	type fields struct {
		words []uint64
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			if got := s.Has(tt.args.x); got != tt.want {
				t.Errorf("IntSet.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}
