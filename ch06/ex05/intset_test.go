package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func Example_one() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestIntSet_Len(t *testing.T) {

	tests := []struct {
		name string
		adds []int
		want int
	}{
		{"no elems", []int{}, 0},
		{"one elems", []int{1}, 1},
		{"two elems", []int{1, 2}, 2},
		{"dup elems", []int{1, 64, 64}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var x IntSet
			for _, a := range tt.adds {
				x.Add(a)
			}
			if got := x.Len(); got != tt.want {
				t.Errorf("x.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Remove(t *testing.T) {
	t.Run("remove exist element", func(t *testing.T) {
		// setup
		var x IntSet
		x.Add(1)
		x.Add(3)

		// execute
		x.Remove(1)

		// assert
		if x.Has(1) {
			t.Errorf("After x.Remove(1), x.Has(1) is true, x: %s", x.String())
		}
		if !x.Has(3) {
			t.Errorf("After x.Remove(1), x.Has(3) is false, x: %s", x.String())
		}
	})

	t.Run("remove not exist element", func(t *testing.T) {
		// setup
		var x IntSet
		x.Add(1)
		x.Add(3)

		// execute
		x.Remove(2)

		// assert
		if !x.Has(1) {
			t.Errorf("After x.Remove(2), x.Has(1) is false, x: %s", x.String())
		}
		if x.Has(2) {
			t.Errorf("After x.Remove(2), x.Has(2) is true, x: %s", x.String())
		}
	})

	t.Run("remove from no elements", func(t *testing.T) {
		// setup
		var x IntSet

		// execute
		x.Remove(2)

		// assert
		if x.Has(2) {
			t.Errorf("After x.Remove(2), x.Has(2) is true, x: %s", x.String())
		}
	})
}

func TestIntSet_Clear(t *testing.T) {
	t.Run("clear elements", func(t *testing.T) {
		// setup
		var x IntSet
		x.Add(1)
		x.Add(3)

		// execute
		x.Clear()

		// assert
		if x.Len() > 0 {
			t.Errorf("After x.Clear(), x.Len() got %d, want 0", x.Len())
		}
		if x.Has(1) {
			t.Errorf("After x.Clear(), x.Has(1) is true, x: %s", x.String())
		}
		if x.Has(3) {
			t.Errorf("After x.Clear(), x.Has(3) is true, x: %s", x.String())
		}
	})

	t.Run("clear from no elements set", func(t *testing.T) {
		// setup
		var x IntSet

		// execute
		x.Clear()

		// assert
		if x.Len() > 0 {
			t.Errorf("After x.Clear(), x.Len() got %d, want 0", x.Len())
		}

	})
}

func TestIntSet_Copy(t *testing.T) {
	t.Run("copy elements", func(t *testing.T) {
		// setup
		var x IntSet
		x.Add(1)
		x.Add(3)

		// execute
		y := *x.Copy()
		x.Add(4)
		x.Remove(1)

		// assert
		if y.Len() != 2 {
			t.Errorf("y.Len() got %d, want 2", y.Len())
		}
		if !y.Has(1) {
			t.Errorf("y.Has(1) is false, y: %s", y.String())
		}
		if !y.Has(3) {
			t.Errorf("y.Has(3) is false, y: %s", y.String())
		}
		if y.Has(4) {
			t.Errorf("y.Has(4) is true, y: %s", y.String())
		}
	})

	t.Run("copy no elements set", func(t *testing.T) {
		// setup
		var x IntSet

		// execute
		y := *x.Copy()
		x.Add(4)
		x.Remove(1)

		// assert
		if y.Len() != 0 {
			t.Errorf("y.Len() got %d, want 0", y.Len())
		}
		if y.Has(4) {
			t.Errorf("y.Has(4) is true, y: %s", y.String())
		}
	})

}

func TestIntSet_AddAll(t *testing.T) {
	// setup
	var x IntSet

	// execute
	x.AddAll(1, 2, 64)

	// assert
	if x.Len() != 3 {
		t.Errorf("x.Len() got %d, want 3", x.Len())
	}
	if !x.Has(1) {
		t.Errorf("x.Has(1) is false, x: %s", x.String())
	}
	if !x.Has(2) {
		t.Errorf("x.Has(2) is false, x: %s", x.String())
	}
	if !x.Has(64) {
		t.Errorf("x.Has(64) is false, x: %s", x.String())
	}
}

func TestIntSet_IntersectWith(t *testing.T) {
	tests := []struct {
		name  string
		xAdds []int
		yAdds []int
		want  []int
	}{
		{"1 same element", []int{1, 2, 3}, []int{2, 4, 6}, []int{2}},
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
