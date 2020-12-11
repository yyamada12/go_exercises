package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	targets, err := loadTargets(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}

	dependedBy, err := loadDependedBy()
	if err != nil {
		log.Fatal(err)
	}

	printDependsOnTargets(targets, dependedBy)
}

func printUsage() {
	fmt.Println(`USAGE:
	go run deps.go [PACKAGES]
	
PACKAGES:
	target packages like "strconv" "strings" or "...xml..."`)
}

type golistResult struct {
	ImportPath string
	Deps       []string
}

func loadTargets(targets ...string) ([]string, error) {
	res := []string{}
	results, err := execGoList(targets...)
	if err != nil {
		return nil, err
	}

	for _, r := range results {
		res = append(res, r.ImportPath)
	}
	return res, nil
}

func loadDependedBy() (map[string][]string, error) {
	results, err := execGoList("...")
	if err != nil {
		return nil, err
	}

	// The key is depended on by values
	dependedBy := map[string][]string{}
	for _, r := range results {
		for _, dep := range r.Deps {
			dependedBy[dep] = append(dependedBy[dep], r.ImportPath)
		}
	}
	return dependedBy, nil

}

func execGoList(targets ...string) ([]golistResult, error) {
	results := []golistResult{}

	args := append([]string{"list", "-json"}, targets...)
	out, _ := exec.Command("go", args...).Output()

	decoder := json.NewDecoder(bytes.NewBuffer(out))
	for decoder.More() {
		res := golistResult{}
		err := decoder.Decode(&res)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func printDependsOnTargets(targets []string, dependedBy map[string][]string) {
	seen := map[string]bool{}
	worklist := targets
	for len(worklist) > 0 {
		crt := worklist[0]
		worklist = worklist[1:]
		if seen[crt] {
			continue
		}
		seen[crt] = true
		worklist = append(worklist, dependedBy[crt]...)
		fmt.Println(crt)
	}
}
