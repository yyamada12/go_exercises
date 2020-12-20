package intset

import (
	"testing"
)

func Test_wordsAndMap(t *testing.T) {
	var x, y IntSet
	x2 := NewMapIntSet()
	y2 := NewMapIntSet()
	x.Add(1)
	x2.Add(1)
	x.Add(144)
	x2.Add(144)
	x.Add(9)
	x2.Add(9)
	if x.String() != x2.String() {
		t.Errorf("x got %s, want %s", x.String(), x2.String())
	}

	y.Add(9)
	y2.Add(9)
	y.Add(42)
	y2.Add(42)

	x.UnionWith(&y)
	x2.UnionWith(y2)
	if x.String() != x2.String() {
		t.Errorf("x union with y got %s, want %s", x.String(), x2.String())
	}

	if x.Has(9) != x2.Has(9) {
		t.Errorf("x.Has(9) got %t, want %t", x.Has(9), x2.Has(9))
	}
	if x.Has(123) != x2.Has(123) {
		t.Errorf("x.Has(123) got %t, want %t", x.Has(123), x2.Has(123))
	}
}
