package intset

import (
	"reflect"
	"testing"
)

func TestIntSet_Elems(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"no elements", []int{}},
		{"1 element", []int{1}},
		{"2 elements", []int{1, 64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var x IntSet
			x.AddAll(tt.want...)
			got := x.Elems()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("x.Elems() got %d, want %d", got, tt.want)
				t.Errorf("x: %s", x.String())
			}
		})
	}
}
