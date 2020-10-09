package main

import "testing"

var noloop = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var depthEqualN = map[string][]string{
	"data structures": {"discrete math"},
	"discrete math":   {"intro to programming"},
}

func Test_toposort(t *testing.T) {
	t.Run("with loop", func(tt *testing.T) {
		_, ok := topoSort(prereqs)
		if ok {
			tt.Errorf("loop should be detected")
		}
	})

	t.Run("no loop", func(tt *testing.T) {
		got, ok := topoSort(noloop)
		if !ok {
			tt.Errorf("loop should not be detected")
		}
		for i, course := range got {
			for _, dep := range noloop[course] {
				if !contains(got[:i], dep) {
					t.Errorf("%s must be in front of %s, got: %q", dep, course, got)
				}
			}
		}
	})

	t.Run("depth == item num", func(tt *testing.T) {
		got, ok := topoSort(depthEqualN)
		if !ok {
			tt.Errorf("loop should not be detected")
		}
		for i, course := range got {
			for _, dep := range prereqs[course] {
				if !contains(got[:i], dep) {
					t.Errorf("%s must be in front of %s, got: %q", dep, course, got)
				}
			}
		}
	})
}

func contains(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}
