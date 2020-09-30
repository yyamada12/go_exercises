// Rev reverses a slice.
package main

import "testing"

func Test_reverse(t *testing.T) {
	// setup
	s := [6]int{0, 1, 2, 3, 4, 5}

	// exec
	reverse(&s)

	// assert
	want := [6]int{5, 4, 3, 2, 1, 0}
	if s != want {
		t.Errorf("reverse([0,1,2,3,4,5] = %q, want %q", s, want)
	}
}
