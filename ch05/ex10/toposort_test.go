package main

import "testing"

func Test_toposort(t *testing.T) {
	// execute
	got := topoSort(prereqs)

	// assert
	for i, course := range got {
		for _, dep := range prereqs[course] {
			if !contains(got[:i], dep) {
				t.Errorf("%s must be in front of %s, got: %q", dep, course, got)
			}
		}
	}
}

func contains(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}
