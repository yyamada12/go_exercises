package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(`USAGE:
		go run anagram.go arg1 arg2`)
	}
	s1, s2 := os.Args[1], os.Args[2]
	if isAnagram(s1, s2) {
		fmt.Println("angram")
	} else {
		fmt.Println("not angram")
	}
}

func isAnagram(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]int)
	for _, c := range s1 {
		m[c]++
	}
	for _, c := range s2 {
		m[c]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
