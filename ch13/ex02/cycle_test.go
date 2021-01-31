package cycle

import (
	"testing"
)

func TestCycle(t *testing.T) {
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	type loop struct {
		child *loop
	}

	l := &loop{}
	l.child = l

	s := "hoge"

	for _, test := range []struct {
		v    interface{}
		want bool
	}{
		{a, true},
		{l, true},
		{[]*string{&s, &s}, false},
		{map[string]*string{"a": &s, "b": &s}, false},
	} {
		if HasCycle(test.v) != test.want {
			t.Errorf("HasCycle(%v) = %t",
				test.v, !test.want)
		}
	}
}
