package main

import "fmt"

func main() {
	fmt.Println(uniq([]string{"hoge", "fuga", "fuga", "piyo", "hoge"}))
}

func uniq(s []string) []string {
	if len(s) == 0 {
		return s
	}
	i := 0
	for _, str := range s {
		if s[i] != str {
			i++
			s[i] = str
		}
	}
	return s[:i+1]
}
