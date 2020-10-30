// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"

	"gopl.io/ch2/popcount"
)

const bits = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bits, uint(x%bits)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bits, uint(x%bits)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds the non-negative values xs to the set
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	if len(t.words) < len(s.words) {
		s.words = s.words[:len(t.words)]
	}
}

// DifferenceWith sets s to the intersect of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, w := range t.words {
		if len(s.words) > i {
			s.words[i] &^= w
		}
	}
}

// SymmetricDifference sets s to the intersect of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, w := range t.words {
		if len(s.words) > i {
			s.words[i] ^= w
		} else {
			s.words = append(s.words, w)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bits*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elemnts in the set
func (s *IntSet) Len() int {
	res := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		res += popcount.PopCount(uint64(word))
	}
	return res
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/bits, uint(x%bits)
	if word < len(s.words) {
		s.words[word] &^= (1 << bit)
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = []uint{}
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	res := &IntSet{make([]uint, len(s.words))}
	copy(res.words, s.words)
	return res
}

// Elems returns all elements of the set
func (s *IntSet) Elems() []int {
	res := make([]int, 0, s.Len())
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, bits*i+j)
			}
		}
	}
	return res
}
