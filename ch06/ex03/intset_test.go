package intset

import "testing"

func TestIntSet_IntersectWith(t *testing.T) {
	tests := []struct {
		name  string
		xAdds []int
		yAdds []int
		want  []int
	}{
		{"1 same element, len(s) < len(t)", []int{1, 2, 3}, []int{2, 4, 64}, []int{2}},
		{"1 same element, len(s) > len(t)", []int{1, 2, 65}, []int{2, 4, 6}, []int{2}},
		{"no smage elements", []int{1, 2, 3}, []int{4, 5, 6}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var x, y IntSet
			x.AddAll(tt.xAdds...)
			y.AddAll(tt.yAdds...)
			x.IntersectWith(&y)
			if x.Len() != len(tt.want) {
				t.Errorf("%d intercect with %d has %d elements, want %d", tt.xAdds, tt.yAdds, x.Len(), len(tt.want))
			}
			for _, w := range tt.want {
				if !x.Has(w) {
					t.Errorf("%d intercect with %d got %s, does not have %d", tt.xAdds, tt.yAdds, x.String(), w)
				}
			}
		})
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	tests := []struct {
		name  string
		xAdds []int
		yAdds []int
		want  []int
	}{
		{"all same element", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"1 same element", []int{1, 2, 3}, []int{2, 4, 64}, []int{1, 3}},
		{"no smage elements", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var x, y IntSet
			x.AddAll(tt.xAdds...)
			y.AddAll(tt.yAdds...)
			x.DifferenceWith(&y)
			if x.Len() != len(tt.want) {
				t.Errorf("%d difference with %d has %d elements, want %d", tt.xAdds, tt.yAdds, x.Len(), len(tt.want))
			}
			for _, w := range tt.want {
				if !x.Has(w) {
					t.Errorf("%d difference with %d got %s, does not have %d", tt.xAdds, tt.yAdds, x.String(), w)
				}
			}
		})
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	tests := []struct {
		name  string
		xAdds []int
		yAdds []int
		want  []int
	}{
		{"all same element", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"1 same element", []int{1, 2, 3}, []int{2, 4, 64}, []int{1, 3, 4, 64}},
		{"no smage elements", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var x, y IntSet
			x.AddAll(tt.xAdds...)
			y.AddAll(tt.yAdds...)
			x.SymmetricDifference(&y)
			if x.Len() != len(tt.want) {
				t.Errorf("%d symmetric difference with %d has %d elements, want %d", tt.xAdds, tt.yAdds, x.Len(), len(tt.want))
			}
			for _, w := range tt.want {
				if !x.Has(w) {
					t.Errorf("%d symmetric difference with %d got %s, does not have %d", tt.xAdds, tt.yAdds, x.String(), w)
				}
			}
		})
	}
}
