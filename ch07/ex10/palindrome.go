package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "mom"
	if isPalindrome(sortableString(s)) {
		fmt.Println(s, "is a palindrome")
	}
}

func isPalindrome(s sort.Interface) bool {
	l := s.Len()
	for i := 0; i*2 < l; i++ {
		if s.Less(i, l-1-i) != s.Less(l-1-i, i) {
			return false
		}
	}
	return true
}

type sortableString []rune

func (s sortableString) Len() int {
	return len(s)
}

func (s sortableString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
