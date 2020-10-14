package intset

import (
	"fmt"
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
