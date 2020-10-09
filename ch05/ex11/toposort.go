// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
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
	"linear algebra":        {"calculus"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	_, ok := topoSort(prereqs)
	fmt.Println(ok)
}

func topoSort(m map[string][]string) ([]string, bool) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	allItems := make(map[string]bool)
	for item, deps := range m {
		allItems[item] = true
		for _, dep := range deps {
			allItems[dep] = true
		}
	}
	var depth int
	n := len(allItems)

	visitAll = func(items []string) {
		depth++
		if depth > n+1 {
			return
		}
		for _, item := range items {
			visitAll(m[item])
			if !seen[item] {
				seen[item] = true
				order = append(order, item)
			}
		}
		if depth > n+1 {
			return
		}
		depth--
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	if depth > n {
		return nil, false
	} else {
		return order, true
	}
}
