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

	targets, err := loadTargets(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	deps, err := loadDeps()
	if err != nil {
		log.Fatal(err)
	}

	printTargetDeps(targets, deps)
}

func printUsage() {
	fmt.Println(`USAGE:
	go run deps.go PACKAGE
	
PACKAGE:
	target package like "strconv"`)
}

type listResult struct {
	Name string
	Deps []string
}

func loadTargets(target string) ([]string, error) {
	out, err := exec.Command("go", "list", "-json", target).Output()
	if err != nil {
		return nil, err
	}
	var jsonResult listResult
	err = json.Unmarshal(out, &jsonResult)
	if err != nil {
		return nil, err
	}
	return jsonResult.Deps, nil
}

func loadDeps() (map[string][]string, error) {
	out, _ := exec.Command("go", "list", "-json", "...").Output()
	decoder := json.NewDecoder(bytes.NewBuffer(out))

	deps := map[string][]string{}
	for decoder.More() {
		var jsonResult listResult
		decoder.Decode(&jsonResult)
		deps[jsonResult.Name] = jsonResult.Deps
	}
	return deps, nil
}

func printTargetDeps(targets []string, deps map[string][]string) {
	seen := map[string]bool{}
	worklist := targets
	for len(worklist) > 0 {
		crt := worklist[0]
		worklist = worklist[1:]
		if seen[crt] {
			continue
		}
		seen[crt] = true
		worklist = append(worklist, deps[crt]...)
		fmt.Println(crt)
	}
}
