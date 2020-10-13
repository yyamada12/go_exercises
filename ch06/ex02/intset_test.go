package intset

import (
	"testing"
)

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
