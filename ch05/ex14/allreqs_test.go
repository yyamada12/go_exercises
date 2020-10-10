// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_breadthFirst(t *testing.T) {
	// execute
	var got []string
	breadthFirst(func(item string) []string {
		got = append(got, item)
		return prereqs[item]
	}, prereqs["databases"])

	// assert
	want := []string{"data structures", "discrete math", "intro to programming"}
	sort.Strings(got)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}

}
