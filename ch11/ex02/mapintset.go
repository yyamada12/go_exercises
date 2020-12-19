// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"strconv"
	"strings"
)

// An MapIntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type MapIntSet struct {
	m map[int]struct{}
}

// NewMapIntSet initializes map and returns *MapIntSet
func NewMapIntSet() *MapIntSet {
	return &MapIntSet{map[int]struct{}{}}
}

// Has reports whether the set contains the non-negative value x.
func (s *MapIntSet) Has(x int) bool {
	_, ok := s.m[x]
	return ok
}

// Add adds the non-negative value x to the set.
func (s *MapIntSet) Add(x int) {
	s.m[x] = struct{}{}
}

// UnionWith sets s to the union of s and t.
func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for elem := range t.m {
		s.m[elem] = struct{}{}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *MapIntSet) String() string {
	elems := []string{}
	for elem := range s.m {
		elems = append(elems, strconv.Itoa(elem))
	}

	return "{" + strings.Join(elems, " ") + "}"
}
