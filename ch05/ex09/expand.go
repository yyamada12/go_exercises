package main

import (
	"regexp"
	"strings"
)

func main() {
	println(expand("$echo", echo))
}

func expand(s string, f func(string) string) string {
	r := regexp.MustCompile(`\$[\w]+`)
	matches := r.FindAllString(s, -1)
	for _, match := range matches {
		s = strings.Replace(s, match, f(match[1:]), 1)
	}
	return s
}

func echo(s string) string {
	return s
}
